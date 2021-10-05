package session

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/marc-campbell/nicedishy/pkg/session/types"
	"github.com/marc-campbell/nicedishy/pkg/stores"
	usertypes "github.com/marc-campbell/nicedishy/pkg/user/types"
	"github.com/pkg/errors"
)

var ErrExpiredSession = errors.New("expired session")

func IsExpiredSession(err error) bool {
	if err == nil {
		return false
	}

	return err == ErrExpiredSession
}

func CreateSessionForUser(ctx context.Context, user *usertypes.User, accessToken string) (*types.Session, error) {
	sess, err := stores.GetStore().CreateSession(ctx, user, accessToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create session")
	}

	return sess, nil
}

func Token(ctx context.Context, sess *types.Session) (string, error) {
	user, err := stores.GetStore().GetUserByID(ctx, sess.UserID)
	if err != nil {
		return "", errors.Wrap(err, "failed to get user")
	}

	claims := jwt.MapClaims{
		"exp":   sess.ExpiresAt.Unix(),
		"nbf":   time.Now().Unix(),
		"aud":   "nicedishy",
		"iss":   "nicedishy-api",
		"sub":   sess.ID,
		"email": user.EmailAddress,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SESSION_KEY")))
	if err != nil {
		return "", errors.Wrap(err, "failed to sign jwt")
	}

	return fmt.Sprintf("Bearer %s", signedToken), nil
}

func Parse(signedToken string) (*types.Session, error) {
	if signedToken == "" {
		return nil, errors.New("missing token")
	}
	tokenParts := strings.Split(signedToken, " ")
	if len(tokenParts) != 2 {
		return nil, fmt.Errorf("invalid number of components in authorization header, expected 2, received %d", len(tokenParts))
	}
	if tokenParts[0] != "Bearer" {
		return nil, errors.New("expected bearer token")
	}

	token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SESSION_KEY")), nil
	})

	if err != nil {
		if err.Error() == "Token is expired" {
			return nil, ErrExpiredSession
		}

		return nil, errors.Wrap(err, "failed to parse jwt token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return stores.GetStore().GetSession(context.TODO(), claims["sub"].(string))
	}

	return nil, errors.New("not a valid jwt")
}
