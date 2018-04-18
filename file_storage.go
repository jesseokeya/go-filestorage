package filestorage

// Cache holds logic for file_storage interaction
type Cache struct {
	store store
}

// Connect connects o the file_storage database instance
func Connect() Cache {
	path := "database/file_storage.json"
	c := Cache{store: store{"", path}}
	c.store.init()
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
