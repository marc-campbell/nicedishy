package pgstore

import (
	"context"
	"fmt"

	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

func (s PGStore) ListDishies(ctx context.Context, userID string) ([]*dishytypes.Dishy, error) {
	pg := persistence.MustGetPGSession()

	query := `select id, created_at, name from dishy where user_id = $1`
	rows, err := pg.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying for dishies: %v", err)
	}

	dishies := []*dishytypes.Dishy{}
	for rows.Next() {
		var d dishytypes.Dishy
		err := rows.Scan(&d.ID, &d.CreatedAt, &d.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning for dishies: %v", err)
		}
		dishies = append(dishies, &d)
	}

	return dishies, nil
}
