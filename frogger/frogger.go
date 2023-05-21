// jsondb - a tiny JSON database
// Copyright (c) 2019, 2023 Steve Domino, Michael D Henderson

// Package frogger implements a very simple logger.
package frogger

import "log"

type Logger struct{}

// Debug implements the jsondb.Logger interface.
func (l *Logger) Debug(format string, args ...any) {
	log.Printf(format, args...)
}

// Error implements the jsondb.Logger interface.
func (l *Logger) Error(format string, args ...any) {
	log.Printf(format, args...)
}

// Fatal implements the jsondb.Logger interface.
func (l *Logger) Fatal(format string, args ...any) {
	log.Fatalf(format, args...)
}

// Info implements the jsondb.Logger interface.
func (l *Logger) Info(format string, args ...any) {
	log.Printf(format, args...)
}

// Trace implements the jsondb.Logger interface.
func (l *Logger) Trace(format string, args ...any) {
	log.Printf(format, args...)
}

// Warn implements the jsondb.Logger interface.
func (l *Logger) Warn(format string, args ...any) {
	log.Printf(format, args...)
}
