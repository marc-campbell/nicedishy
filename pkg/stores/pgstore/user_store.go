package pgstore

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
	usertypes "github.com/marc-campbell/nicedishy/pkg/user/types"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
)

func (s PGStore) CreateUserToken(ctx context.Context, userID string) (string, error) {
	pg := persistence.MustGetPGSession()

	r, err := ksuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "create random")
	}
	token := r.String()
	sha := fmt.Sprintf("%x", sha256.Sum256([]byte(token)))
	query := `insert into user_token (token_sha, created_at, user_id) values ($1, $2, $3)`
	_, err = pg.Exec(ctx, query, sha, time.Now(), userID)
	if err != nil {
		return "", errors.Wrap(err, "insert token sha")
	}

	return token, nil
}

func (s PGStore) GetUserByEmail(ctx context.Context, email string) (*usertypes.User, error) {
	pg := persistence.MustGetPGSession()

	query := `select id from google_user where email_address = $1`
	row := pg.QueryRow(ctx, query, email)

	id := ""
	if err := row.Scan(&id); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, errors.Wrap(err, "scan user id")
	}

	return s.GetUserByID(ctx, id)
}

func (s PGStore) GetUserByID(ctx context.Context, id string) (*usertypes.User, error) {
	pg := persistence.MustGetPGSession()

	query := `select id, email_address, avatar_url, created_at, last_login_at, is_waitlisted from google_user where id = $1`
	row := pg.QueryRow(ctx, query, id)

	user := usertypes.User{}
	var lastLoginAt sql.NullTime
	if err := row.Scan(&user.ID, &user.EmailAddress, &user.AvatarURL, &user.CreatedAt, &lastLoginAt, &user.IsWaitlisted); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, errors.Wrap(err, "scan user row")
	}

	if lastLoginAt.Valid {
		user.LastLoginAt = &lastLoginAt.Time
	}

	return &user, nil
}

func (s PGStore) CreateUser(ctx context.Context, emailAddress string, avatarURL string) (*usertypes.User, error) {
	pg := persistence.MustGetPGSession()

	id, err := ksuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate user id")
	}

	now := time.Now()

	user := usertypes.User{
		ID:           id.String(),
		EmailAddress: emailAddress,
		AvatarURL:    avatarURL,
		CreatedAt:    now,
		LastLoginAt:  &now,
		IsWaitlisted: true,
	}

	query := `insert into google_user (id, email_address, avatar_url, created_at, last_login_at, is_waitlisted) values ($1, $2, $3, $4, $5, $6)`
	_, err = pg.Exec(ctx, query, user.ID, user.EmailAddress, user.AvatarURL, user.CreatedAt, user.LastLoginAt, user.IsWaitlisted)
	if err != nil {
		return nil, errors.Wrap(err, "insert user")
	}

	// because uses are hard coded to be on the waitlist, we also insert this
	query = `insert into google_user_waitlist (id, created_at, already_have_dishy, how_long_with_dishy, primary_or_backup, operating_systems, why_access)
		values ($1, $2, $3, $4, $5, $6, $7)`
	_, err = pg.Exec(ctx, query, user.ID, user.CreatedAt, "", "", "", []string{}, []string{})
	if err != nil {
		return nil, errors.Wrap(err, "insert user waitlist")
	}

	return &user, nil
}
