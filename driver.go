// jsondb - a tiny JSON database
// Copyright (c) 2019, 2023 Steve Domino, Michael D Henderson

package jsondb

import (
	"fmt"
	"github.com/maloquacious/jsondb/frogger"
	"os"
	"path/filepath"
	"sync"
)

// New creates a new database at the desired directory location, and
// returns a *DB to then use for interacting with the database.
func New(dir string, options ...Option) (*DB, error) {
	db := DB{
		dir:     filepath.Clean(dir),
		log:     &frogger.Logger{}, // default logger
		mutexes: make(map[string]*sync.Mutex),
	}

	for _, opt := range options {
		if err := opt(&db); err != nil {
			return nil, err
		}
	}

	// if the database already exists, just use it
	if _, err := os.Stat(dir); err == nil {
		db.log.Debug("jsondb: using '%s' (database already exists)\n", dir)
		return &db, ErrExists
	}

	// if the database doesn't exist create it
	db.log.Debug("jsondb: creating database at '%s'...\n", dir)
	return &db, os.MkdirAll(dir, 0755)
}

// Delete locks the database then attempts to remove the collection/resource
// specified by [path]
func (db *DB) Delete(collection, resource string) error {
	mutex := db.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	path := filepath.Join(collection, resource)
	dir := filepath.Join(db.dir, path)

	if fi, err := stat(dir); err != nil {
		return fmt.Errorf("%q: %w\n", path, err)
	} else if fi == nil {
		return fmt.Errorf("%q: %w\n", path, os.ErrNotExist)
	} else if fi.Mode().IsDir() {
		// remove directory and all contents
		return os.RemoveAll(dir)
	} else if fi.Mode().IsRegular() {
		// remove file
		return os.RemoveAll(dir + ".json")
	}

	return nil
}

// Read a record from the database
func (db *DB) Read(collection, resource string, v any) error {
	// ensure there is a place to save record
	if collection == "" {
		return ErrMissingCollection
	}
	// ensure there is a resource (name) to save record as
	if resource == "" {
		return ErrMissingResource
	}

	//
	record := filepath.Join(db.dir, collection, resource)

	// read record from database; if the file doesn't exist `read` will return an error
	return read(record, v)
}

// ReadAll records from a collection; this is returned as a slice of strings because
// there is no way of knowing what type the record is.
func (db *DB) ReadAll(collection string) ([][]byte, error) {

	// ensure there is a collection to read
	if collection == "" {
		return nil, ErrMissingCollection
	}

	//
	dir := filepath.Join(db.dir, collection)

	// read all the files in the transaction.Collection; an error here just means
	// the collection is either empty or doesn't exist
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return readAll(files, dir)
}

// Write locks the database and attempts to write the record to the database under
// the [collection] specified with the [resource] name given.
func (db *DB) Write(collection, resource string, v any) error {
	// ensure there is a place to save record
	if collection == "" {
		return ErrMissingCollection
	}
	// ensure there is a resource (name) to save record as
	if resource == "" {
		return ErrMissingResource
	}

	mutex := db.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	//
	dir := filepath.Join(db.dir, collection)
	fnlPath := filepath.Join(dir, resource+".json")
	tmpPath := fnlPath + ".tmp"

	return write(dir, tmpPath, fnlPath, v)
}

// getOrCreateMutex creates a new collection specific mutex any time a collection
// is being modified to avoid unsafe operations
func (db *DB) getOrCreateMutex(collection string) *sync.Mutex {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	m, ok := db.mutexes[collection]

	// if the mutex doesn't exist make it
	if !ok {
		m = &sync.Mutex{}
		db.mutexes[collection] = m
	}

	return m
}
