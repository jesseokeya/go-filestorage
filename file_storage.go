package filestorage

// Cache holds logic for file_storage interaction
type Cache struct {
	store store
}

// Connect connects o the file_storage database instance
func Connect() Cache {
	path := "database/file_storage.json"
	return Cache{store: store{"", path}}
}

// Name the database any given name you wish
func (c *Cache) Name(db string) {
	c.store.Name = db
}
