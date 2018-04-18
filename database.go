package filestorage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Store tracks all needed data structures for database instance
type store struct {
	Path string
	Name string
}

// init creates the json path to store data
func (s *store) init() {
	s.createFileStore()
}

// createFileStore creates a json file where data would be stored in memory
func (s *store) createFileStore() {
	newpath := filepath.Join(".", "database")
	os.MkdirAll(newpath, os.ModePerm)

	_, err := os.Stat(s.Path)
	if os.IsNotExist(err) {
		if len(s.Path) == 0 {
			s.Path = "database/storage.json"
		}
		file, err := os.Create(s.Path)
		handleError(err)
		emptyData := make([]interface{}, 0)
		data, err := json.Marshal(emptyData)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(s.Path, data, 0777)
		handleError(err)
		defer file.Close()
	}
}

// writeToFile writes the data to be stored in a location in memory
func (s *store) writeToFile(p interface{}) {
	if len(s.readFromFile()) == 0 {
		result := make([]interface{}, 0)
		result = append(result, p)
		data, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(s.Path, data, 0777)
		handleError(err)
	} else {
		previous := s.readFromFile()

		result := make([]interface{}, 0)
		for _, data := range previous {
			fmt.Println(data)
			result = append(result, data)
		}
		result = append(result, p)
		data, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(s.Path, data, 0777)
		handleError(err)
	}
}

// GetAll returns a byte array of all data from the file
func (s *store) readFromFile() []interface{} {
	var result []interface{}
	data, err := ioutil.ReadFile(s.Path)
	if err != nil {
		handleError(err)
	}
	err = json.Unmarshal(data, &result)
	handleError(err)
	return result
}

// delete removes file path and file where data was stored
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
