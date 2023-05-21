// jsondb - a tiny JSON database
// Copyright (c) 2019, 2023 Steve Domino, Michael D Henderson

// Package cerrors implements constant error wrappers from Cheney.
//
// Usage is something like:
//
//	const ErrFoo = cerrors.Error("foo")
//	if errors.Is(err, ErrFoo) { ... }
package cerrors

// The original article by Cheney is at:
//   https://dave.cheney.net/2016/04/07/constant-errors
// Myren has a good article:
//   https://smyrman.medium.com/writing-constant-errors-with-go-1-13-10c4191617

// Error allows us to create constant error values
type Error string

// Error implements the Error interface.
func (err Error) Error() string {
	return string(err)
}
