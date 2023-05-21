// jsondb - a tiny JSON database
// Copyright (c) 2019, 2023 Steve Domino, Michael D Henderson

// Package jsondb implements a fork of Steve Domino's Scribble.
package jsondb

import "sync"

// Version is the current version of the project
const Version = "2.0.0"

// DB is what is used to interact with the JSONDB database.
// It runs transactions, and provides log output.
type DB struct {
	mutex   sync.Mutex
	mutexes map[string]*sync.Mutex
	dir     string // the directory where JSONDB will create the database
	log     Logger // the logger JSONDB will log to
}
