package stores

import (
	"context"

	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
	sessiontypes "github.com/marc-campbell/nicedishy/pkg/session/types"
	"github.com/marc-campbell/nicedishy/pkg/stores/pgstore"
	tokentypes "github.com/marc-campbell/nicedishy/pkg/token/types"
	usertypes "github.com/marc-campbell/nicedishy/pkg/user/types"
)

var _ Store = (*pgstore.PGStore)(nil)

var (
	hasStore    = false
	globalStore Store
)

func GetStore() Store {
	if !hasStore {
		globalStore = pgstore.PGStore{}
	}

	return globalStore
}

type Store interface {
	IsNotFound(err error) bool

	CreateSession(ctx context.Context, user *usertypes.User, accessToken string) (*sessiontypes.Session, error)
	GetSession(ctx context.Context, id string) (*sessiontypes.Session, error)

	GetUserByEmail(ctx context.Context, emailAddress string) (*usertypes.User, error)
	GetUserByID(ctx context.Context, id string) (*usertypes.User, error)
	CreateUser(ctx context.Context, emailAddress string, avatarURL string) (*usertypes.User, error)
	CreateUserToken(ctx context.Context, userID string) (string, error)
	GetToken(ctx context.Context, token string) (*tokentypes.Token, error)

	ListDishies(ctx context.Context, userID string) ([]*dishytypes.Dishy, error)
}
