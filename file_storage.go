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
	path := "database/storage.json"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		c := Cache{store: store{"", path}}
		c.store.init()
		return c
	}
	return Cache{store: store{"", path}}
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

// Insert adds an array of new object interface to the database
func (c *Cache) Insert(p ...interface{}) error {
	allData := c.store.readFromFile()
	for _, obj := range p {
		if !c.contains(obj, allData) {
			c.store.writeToFile(obj)
		}
	}
	return nil
}

// InsertOne adds a new object interface to the database
func (c *Cache) InsertOne(p interface{}) error {
	allData := c.store.readFromFile()
	if !c.contains(p, allData) {
		c.store.writeToFile(p)
	}
	return nil
}

func (c *Cache) contains(i interface{}, list ...interface{}) bool {
	for _, item := range list {
		if item == i {
			return true
		}
	}
	return false
}
