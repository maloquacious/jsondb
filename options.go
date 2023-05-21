// jsondb - a tiny JSON database
// Copyright (c) 2019, 2023 Steve Domino, Michael D Henderson

package jsondb

type Option func(db *DB) error

func SetLogger(l Logger) func(*DB) error {
	return func(db *DB) error {
		db.log = l
		return nil
	}
}
