package types

import (
	dishytypes "github.com/marc-campbell/nicedishy/pkg/dishy/types"
)

type Token struct {
	SHA   string
	Dishy *dishytypes.Dishy
}
