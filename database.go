package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"os"
)

// Store tracks all needed data structures for database instance
type store struct {
	Path string
	Name string
}

// createFileStore creates a json file where data would be stored in memory
func (s *store) createFileStore() {
	_, err := os.Stat(s.Path)
	if os.IsNotExist(err) {
		file, err := os.Create(s.Path)
		handleError(err)
		defer file.Close()
	}
}

// writeToFile writes the data to be stored in a location in memory
func (s *store) writeToFile(p ...interface{}) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	handleError(err)
	err = ioutil.WriteFile(s.Path, buf.Bytes(), 0777)
	handleError(err)
}

// GetAll returns a byte array of all data from the file
func (s *store) readFromFile() []interface{} {
	var result []interface{}
	data, err := ioutil.ReadFile(s.Path)
	if err != nil {
		handleError(err)
	}
	err = json.Unmarshal(data, result)
	handleError(err)
	return result
}

func (s *store) deleteFile(path string) error {
	err := os.Remove(path)
	return err
}

// HandleError handles error that come up during runtime
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
