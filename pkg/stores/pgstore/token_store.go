package pgstore

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/marc-campbell/nicedishy/pkg/persistence"
	tokentypes "github.com/marc-campbell/nicedishy/pkg/token/types"
	"github.com/pkg/errors"
)

func (s PGStore) GetToken(ctx context.Context, token string) (*tokentypes.Token, error) {
	pg := persistence.MustGetPGSession()

	sha := fmt.Sprintf("%x", sha256.Sum256([]byte(token)))

	query := `select user_id from user_token where token_sha = $1`
	row := pg.QueryRow(ctx, query, sha)

	var userID string
	if err := row.Scan(&userID); err != nil {
		return nil, errors.Wrap(err, "scan user")
	}

	user, err := s.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	t := tokentypes.Token{
		SHA:  sha,
		User: user,
	}

	return &t, nil
}
