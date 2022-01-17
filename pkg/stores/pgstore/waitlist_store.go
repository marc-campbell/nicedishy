package pgstore

import (
	"context"

	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

func (s PGStore) UpdateUserWaitlistData(ctx context.Context, id string, alreadyHaveDishy string, howLongWithDishy string, primaryOrBackup string, operatingSystems []string, whyAccess []string) error {
	pg := persistence.MustGetPGSession()

	query := `insert into google_user_waitlist
(id, created_at, already_have_dishy, how_long_with_dishy, primary_or_backup, operating_systems, why_access)
values ($1, now(), $2, $3, $4, $5, $6)
on conflict (id) do update set created_at = now(),
already_have_dishy = $2, how_long_with_dishy = $3,
primary_or_backup = $4, operating_systems = $5, why_access = $6`

	_, err := pg.Exec(ctx, query, id, alreadyHaveDishy, howLongWithDishy, primaryOrBackup, operatingSystems, whyAccess)
	if err != nil {
		return err
	}

	return nil
}
