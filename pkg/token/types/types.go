package types

import (
	usertypes "github.com/marc-campbell/nicedishy/pkg/user/types"
)

type Token struct {
	SHA  string
	User *usertypes.User
}
