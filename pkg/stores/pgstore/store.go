package pgstore

import (
	"errors"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrNotImplemented = errors.New("not implemeted")
)

type PGStore struct {
}

func (s PGStore) IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	return err == ErrNotFound
}
