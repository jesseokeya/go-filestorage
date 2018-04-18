package filestorage

import "os"

// Store tracks all needed data structures for database instance
type Store struct {
	Schema interface{}
	Name   string
}

// Initialize connects to the database but name and scheme is required
func (s *Store) Initialize(i interface{}, n string) {
	s.Schema = i
	s.Name = n
	path := "./database/file_storage.json"
	CreateFileStore(path)
}

// CreateFileStore creates a json file where data would be stored in memory
func CreateFileStore(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		HandleError(err)
		defer file.Close()
	}
}

// HandleError handles error that come up during runtime
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
