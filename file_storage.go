package filestorage

import (
	"os"
)

// Cache holds logic for file_storage interaction
type Cache struct {
	store store
}

// Connect connects o the file_storage database instance
func Connect() Cache {
	c := Cache{}
	c.store.Name = ""
	c.store.Path = "database/storage.json"
	if _, err := os.Stat(c.store.Path); os.IsNotExist(err) {
		c.store.init()
		return c
	}
	return c
}

// Name the database any given name you wish
func (c *Cache) Name(db string) {
	c.store.Name = db
}

// GetPath returns the path to the storage in memory
func (c *Cache) GetPath() string {
	return c.store.Path
}

// GetName returns the database name
func (c *Cache) GetName() string {
	return c.store.Name
}

// FindAll returns all data in the database
func (c *Cache) FindAll() []interface{} {
	return c.store.readFromFile()
}

// InsertOne adds a new object interface to the database
func (c *Cache) InsertOne(p interface{}) error {
	allData := c.store.readFromFile()
	if !c.contains(p, allData) {
		c.store.writeToFile(p)
	}
	return nil
}

// contains ensures integrity of the database and eliminates repeteation
func (c *Cache) contains(i interface{}, list ...interface{}) bool {
	for _, item := range list {
		if item == i {
			return true
		}
	}
	return false
}
