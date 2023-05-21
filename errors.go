// jsondb - a tiny JSON database
// Copyright (c) 2019, 2023 Steve Domino, Michael D Henderson

package jsondb

import "github.com/maloquacious/jsondb/cerrors"

const (
	ErrExists            = cerrors.Error("db already exists")
	ErrMissingCollection = cerrors.Error("missing collection - no place to save record")
	ErrMissingResource   = cerrors.Error("missing resource - unable to save record")
)
