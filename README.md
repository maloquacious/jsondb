# jsondb
[![GoDoc](https://godoc.org/github.com/boltdb/bolt?status.svg)](http://godoc.org/github.com/maloquacious/jsondb)
[![Go Report Card](https://goreportcard.com/badge/github.com/maloquacious/jsondb)](https://goreportcard.com/report/github.com/maloquacious/jsondb)

A tiny JSON database in Golang.

Forked from [Steve Domino](https://github.com/sdomino)'s [Scribble](https://github.com/sdomino/scribble).

I changed the name of the package because some of my changes are incompatible with the original source.

## Installation

Add `import "github.com/maloquacious/jsondb"` to your file.

## Usage

```go
// a new driver, providing the directory where it will be writing to,
// and a qualified logger if desired
db, err := scribble.New(dir)
if err != nil && !errors.Is(err, jsondb.ErrExists){
  fmt.Println("Error", err)
}

// Write a fish to the database
fish := Fish{}
if err := db.Write("fish", "onefish", fish); err != nil {
  fmt.Println("Error", err)
}

// Read a fish from the database (passing fish by reference)
onefish := Fish{}
if err := db.Read("fish", "onefish", &onefish); err != nil {
  fmt.Println("Error", err)
}

// Read all fish from the database, unmarshaling the response.
records, err := db.ReadAll("fish")
if err != nil {
  fmt.Println("Error", err)
}

fishies := []Fish{}
for _, f := range records {
  fishFound := Fish{}
  if err := json.Unmarshal([]byte(f), &fishFound); err != nil {
    fmt.Println("Error", err)
  }
  fishies = append(fishies, fishFound)
}

// Delete a fish from the database
if err := db.Delete("fish", "onefish"); err != nil {
  fmt.Println("Error", err)
}

// Delete all fish from the database
if err := db.Delete("fish", ""); err != nil {
  fmt.Println("Error", err)
}
```

## Documentation
- Complete documentation is available on [godoc](https://pkg.go.dev/github.com/maloquacious/jsondb).

## TODO
- Support for windows
- Better support for concurrency
- Better support for sub collections
- More methods to allow different types of reads/writes
- More tests (you can never have enough!)
