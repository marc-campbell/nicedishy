package stores

import (
	"context"
	"time"

	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
	mailertypes "github.com/marc-campbell/nicedishy/pkg/mailer/types"
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
	CreateOAuthState(ctx context.Context, next string) (string, error)
	GetOAuthState(ctx context.Context, state string) (bool, string, error)

	GetUserByEmail(ctx context.Context, emailAddress string) (*usertypes.User, error)
	GetUserByID(ctx context.Context, id string) (*usertypes.User, error)
	CreateUser(ctx context.Context, emailAddress string, avatarURL string) (*usertypes.User, error)
	CreateUserToken(ctx context.Context, userID string) (string, error)
	GetToken(ctx context.Context, token string) (*tokentypes.Token, error)
	CreateSessionNonce(ctx context.Context, sessionID string) (string, error)
	GetSessionNonce(ctx context.Context, id string) (string, error)
	GetUserByDishy(ctx context.Context, id string) (*usertypes.User, error)

	ListDishies(ctx context.Context, userID string) ([]*dishytypes.Dishy, error)
	CreateDishy(ctx context.Context, userID string, name string) (*dishytypes.Dishy, error)
	GetDishyForUser(ctx context.Context, id string, userID string) (*dishytypes.Dishy, error)
	CreateDishyToken(ctx context.Context, id string) (string, error)
	GetDishy(ctx context.Context, id string) (*dishytypes.Dishy, error)
	SetDishyLastReceivedStats(ctx context.Context, id string, when time.Time) error
	DeleteDishy(ctx context.Context, id string) error
	UpdateDishyGeo(ctx context.Context, id string, when time.Time, geo *dishytypes.GeoCheck) error

	// stats
	GetConnectedDishyCount(ctx context.Context) (int, error)
	GetTotalDishyCount(ctx context.Context) (int, error)
	GetNewDishyCount(ctx context.Context) (int, error)
	GetHighestDownloadSpeed(ctx context.Context) (float64, error)
	GetAverageDownloadSpeed(ctx context.Context) (float64, error)
	GetLowestPingTime(ctx context.Context) (float64, error)
	GetAveragePingTime(ctx context.Context) (float64, error)

	// one dishy
	GetDishyVersions(ctx context.Context, id string) (string, string, error)

	// mailer
	GetQueuedEmails(ctx context.Context) ([]*mailertypes.Email, error)
	MarkQueuedEmailSent(ctx context.Context, id string) error
	MarkQueuedEmailError(ctx context.Context, id string) error
	QueueEmail(ctx context.Context, fromAddress string, toAddress string, templateID int64, templateContext map[string]interface{}) (*mailertypes.Email, error)
}
