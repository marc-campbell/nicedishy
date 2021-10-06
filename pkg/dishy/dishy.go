package dishy

import (
	"context"
	"fmt"

	"github.com/marc-campbell/nicedishy/pkg/dishy/types"
	"github.com/marc-campbell/nicedishy/pkg/stores"
)

func CreateDefaultForUser(ctx context.Context, userID string) (*types.Dishy, error) {
	// We just create a single dishy named "My Dishy" for the user

	dishy, err := stores.GetStore().CreateDishy(ctx, userID, "My Dishy")
	if err != nil {
		return nil, fmt.Errorf("failed to create default dishy: %w", err)
	}

	return dishy, nil
}
