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

	query := `select dishy_id from dishy_token where token_sha = $1`
	row := pg.QueryRow(ctx, query, sha)

	var dishyID string
	if err := row.Scan(&dishyID); err != nil {
		return nil, errors.Wrap(err, "scan user")
	}

	dishy, err := s.GetDishy(ctx, dishyID)
	if err != nil {
		return nil, errors.Wrap(err, "get dishy")
	}

	t := tokentypes.Token{
		SHA:   sha,
		Dishy: dishy,
	}

	return &t, nil
}
