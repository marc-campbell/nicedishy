package pgstore

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
	sessiontypes "github.com/marc-campbell/nicedishy/pkg/session/types"
	usertypes "github.com/marc-campbell/nicedishy/pkg/user/types"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
)

func (s PGStore) GetSessionNonce(ctx context.Context, id string) (string, error) {
	pg := persistence.MustGetPGSession()

	query := `delete from session_nonce where id = $1 and expire_at > $2 returning session_id`
	row := pg.QueryRow(ctx, query, id, time.Now())
	sessionID := ""
	if err := row.Scan(&sessionID); err != nil {
		if err == pgx.ErrNoRows {
			return "", nil
		}

		return "", errors.Wrap(err, "failed to scan session nonce")
	}

	return sessionID, nil
}

func (s PGStore) CreateSessionNonce(ctx context.Context, sessionID string) (string, error) {
	pg := persistence.MustGetPGSession()

	id, err := ksuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate session id")
	}

	query := `insert into session_nonce (id, expire_at, session_id) values ($1, $2, $3)`
	_, err = pg.Exec(ctx, query, id.String(), time.Now().Add(time.Minute), sessionID)
	if err != nil {
		return "", errors.Wrap(err, "insert session nonce")
	}

	return id.String(), nil
}

func (s PGStore) CreateSession(ctx context.Context, user *usertypes.User, accessToken string) (*sessiontypes.Session, error) {
	pg := persistence.MustGetPGSession()

	id, err := ksuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate session id")
	}

	sess := sessiontypes.Session{
		ID:        id.String(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	query := `insert into session (id, user_id, expire_at, access_token) values ($1, $2, $3, $4)`
	_, err = pg.Exec(ctx, query, sess.ID, sess.UserID, sess.ExpiresAt, accessToken)
	if err != nil {
		return nil, errors.Wrap(err, "insert session")
	}

	return &sess, nil
}

func (s PGStore) GetSession(ctx context.Context, id string) (*sessiontypes.Session, error) {
	pg := persistence.MustGetPGSession()

	query := `select id, user_id, expire_at, access_token from session where id = $1`
	row := pg.QueryRow(ctx, query, id)

	sess := sessiontypes.Session{}
	if err := row.Scan(&sess.ID, &sess.UserID, &sess.ExpiresAt, &sess.AccessToken); err != nil {
		return nil, errors.Wrap(err, "failed to scan session")
	}

	return &sess, nil
}

func (s PGStore) CreateOAuthState(ctx context.Context, next string) (string, error) {
	pg := persistence.MustGetPGSession()

	id, err := ksuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate session id")
	}

	now := time.Now()

	query := `insert into oauth_state (id, created_at, next) values ($1, $2, $3)`
	_, err = pg.Exec(ctx, query, id.String(), now, next)
	if err != nil {
		return "", errors.Wrap(err, "insert oauth state")
	}

	return id.String(), nil
}

func (s PGStore) GetOAuthState(ctx context.Context, id string) (bool, string, error) {
	pg := persistence.MustGetPGSession()

	query := `select id, next from oauth_state where id = $1`
	next := ""

	row := pg.QueryRow(ctx, query, id)
	if err := row.Scan(&id, &next); err != nil {
		if err == pgx.ErrNoRows {
			return false, "", nil
		}

		return false, "", errors.Wrap(err, "failed to scan oauth state")
	}

	query = `delete from oauth_state where id = $1`
	_, err := pg.Exec(ctx, query, id)
	if err != nil {
		return false, "", errors.Wrap(err, "delete oauth state")
	}

	return true, next, nil
}
