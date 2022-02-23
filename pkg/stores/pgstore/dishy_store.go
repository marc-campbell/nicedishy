package pgstore

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
	"github.com/segmentio/ksuid"
)

func (s PGStore) GetDishyVersions(ctx context.Context, id string) (string, string, error) {
	pg := persistence.MustGetMetricsDBSession()

	query := `select software_version, hardware_version from dishy_data where dishy_id = $1
and software_version is not null and hardware_version is not null
order by time desc limit 1`
	row := pg.QueryRow(ctx, query, id)

	softwareVersion := ""
	hardwareVersion := ""

	if err := row.Scan(&softwareVersion, &hardwareVersion); err != nil {
		if err == pgx.ErrNoRows {
			return "", "", nil
		}

		return "", "", fmt.Errorf("error getting versions: %v", err)
	}

	return softwareVersion, hardwareVersion, nil
}

func (s PGStore) ListDishies(ctx context.Context, userID string) ([]*dishytypes.Dishy, error) {
	pg := persistence.MustGetPGSession()

	query := `select id, created_at, last_metric_at, name from dishy where user_id = $1`
	rows, err := pg.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying for dishies: %v", err)
	}

	dishies := []*dishytypes.Dishy{}
	for rows.Next() {
		var d dishytypes.Dishy
		var lastMetricAt sql.NullTime
		err := rows.Scan(&d.ID, &d.CreatedAt, &lastMetricAt, &d.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning for dishies: %v", err)
		}

		if lastMetricAt.Valid {
			d.LastMetricAt = &lastMetricAt.Time
		}

		dishies = append(dishies, &d)
	}

	return dishies, nil
}

func (s PGStore) CreateDishy(ctx context.Context, userID string, name string) (*dishytypes.Dishy, error) {
	pg := persistence.MustGetPGSession()

	id, err := ksuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error creating id: %v", err)
	}

	now := time.Now()
	query := `insert into dishy (id, user_id, created_at, name) values ($1, $2, $3, $4)`
	if _, err := pg.Exec(ctx, query, id.String(), userID, now, name); err != nil {
		return nil, fmt.Errorf("error creating dishy: %v", err)
	}

	return &dishytypes.Dishy{
		ID:        id.String(),
		CreatedAt: now,
		Name:      name,
	}, nil
}

func (s PGStore) GetDishyForUser(ctx context.Context, id string, userID string) (*dishytypes.Dishy, error) {
	pg := persistence.MustGetPGSession()

	query := `select id, created_at, last_metric_at, name from dishy where user_id = $1 and id = $2`
	row := pg.QueryRow(ctx, query, userID, id)

	dishy := dishytypes.Dishy{}
	lastMetricAt := sql.NullTime{}

	if err := row.Scan(&dishy.ID, &dishy.CreatedAt, &lastMetricAt, &dishy.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

	}

	if lastMetricAt.Valid {
		dishy.LastMetricAt = &lastMetricAt.Time
	}

	return &dishy, nil
}

func (s PGStore) CreateDishyToken(ctx context.Context, id string) (string, error) {
	pg := persistence.MustGetPGSession()

	token, err := ksuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("error creating id: %w", err)
	}

	tokenSHA := fmt.Sprintf("%x", sha256.Sum256([]byte(token.String())))
	query := `insert into dishy_token (token_sha, dishy_id) values ($1, $2)`
	if _, err := pg.Exec(ctx, query, tokenSHA, id); err != nil {
		return "", fmt.Errorf("error creating token: %w", err)
	}

	return token.String(), nil
}

func (s PGStore) GetDishy(ctx context.Context, id string) (*dishytypes.Dishy, error) {
	pg := persistence.MustGetPGSession()

	query := `select id, created_at, last_metric_at, last_geocheck_at, name from dishy where id = $1`
	row := pg.QueryRow(ctx, query, id)

	dishy := dishytypes.Dishy{}
	lastMetricAt := sql.NullTime{}
	lastGeocheckAt := sql.NullTime{}

	if err := row.Scan(&dishy.ID, &dishy.CreatedAt, &lastMetricAt, &lastGeocheckAt, &dishy.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

	}

	if lastMetricAt.Valid {
		dishy.LastMetricAt = &lastMetricAt.Time
	}

	if lastGeocheckAt.Valid {
		dishy.LastGeocheckAt = &lastGeocheckAt.Time
	}

	return &dishy, nil
}

func (s PGStore) SetDishyLastReceivedStats(ctx context.Context, id string, when time.Time) error {
	pg := persistence.MustGetPGSession()

	query := `update dishy set last_metric_at = $1 where id = $2 and not exists (select 1 from dishy where id = $3 and last_metric_at > $4)`
	if _, err := pg.Exec(ctx, query, when, id, id, when); err != nil {
		return fmt.Errorf("error setting last_metric_at: %v", err)
	}

	query = `insert into dishy_disconnected_queue (dishy_id, send_at) values ($1, $2) on conflict (dishy_id)
do update set send_at = EXCLUDED.send_at`
	sendAt := time.Now().Add(time.Hour * 6)
	if _, err := pg.Exec(ctx, query, id, sendAt); err != nil {
		return fmt.Errorf("error inserting into dishy_disconnected_queue: %v", err)
	}

	return nil
}

func (s PGStore) DeleteDishy(ctx context.Context, id string) error {
	pg := persistence.MustGetPGSession()

	query := `delete from dishy where id = $1`
	if _, err := pg.Exec(ctx, query, id); err != nil {
		return fmt.Errorf("error deleting dishy: %w", err)
	}

	return nil
}

func (s PGStore) UpdateDishyGeo(ctx context.Context, id string, when time.Time, geo *dishytypes.GeoCheck) error {
	pg := persistence.MustGetMetricsDBSession()

	query := `insert into dishy_geo (time, id, ip_address, continent, country, region, city, org, latitude, longitude) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	if _, err := pg.Exec(ctx, query, when, id, geo.IPAddress, geo.Continent, geo.Country, geo.Region, geo.City, geo.Org, geo.Latitude, geo.Longitude); err != nil {
		return fmt.Errorf("error inserting geo: %w", err)
	}

	otherPg := persistence.MustGetPGSession()
	query = `update dishy set last_geocheck_at = $1 where id = $2`
	if _, err := otherPg.Exec(ctx, query, when, id); err != nil {
		return fmt.Errorf("error updating last_geocheck_at: %w", err)
	}

	return nil
}
