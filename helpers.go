// jsondb - a tiny JSON database
// Copyright (c) 2019, 2023 Steve Domino, Michael D Henderson

package jsondb

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func read(record string, v any) error {
	b, err := os.ReadFile(record + ".json")
	if err != nil {
		return err
	}
	// unmarshal data
	return json.Unmarshal(b, v)
}

func readAll(files []os.DirEntry, dir string) ([][]byte, error) {
	// the files read from the database
	var records [][]byte

	// iterate over each of the files, attempting to read the file. If successful
	// append the files to the collection of read.
	for _, file := range files {
		b, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		// append read file
		records = append(records, b)
	}

	// unmarhsal the read files as a comma delimited byte array
	return records, nil
}

func stat(path string) (fi os.FileInfo, err error) {
	// check for dir, if path isn't a directory check to see if it's a file
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return fi, err
}

func write(dir, tmpPath, dstPath string, v any) error {
	// create collection directory
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// marshal the pointer to a non-struct and indent with tab
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}

	// for a newline on the output
	if len(b) != 0 && b[len(b)-1] != '\n' {
		b = append(b, byte('\n'))
	}

	// write marshaled data to the temp file
	if err := os.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}

	// move final file into place
	return os.Rename(tmpPath, dstPath)
}
