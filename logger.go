// jsondb - a tiny JSON database
// Copyright (c) 2019, 2023 Steve Domino, Michael D Henderson

package jsondb

// Logger is a generic logger interface.
type Logger interface {
	Fatal(format string, args ...any)
	Error(format string, args ...any)
	Warn(format string, args ...any)
	Info(format string, args ...any)
	Debug(format string, args ...any)
	Trace(format string, args ...any)
}
