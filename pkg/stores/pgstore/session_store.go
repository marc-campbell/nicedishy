package pgstore

import (
	"context"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/persistence"
	sessiontypes "github.com/marc-campbell/nicedishy/pkg/session/types"
	usertypes "github.com/marc-campbell/nicedishy/pkg/user/types"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
)

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
