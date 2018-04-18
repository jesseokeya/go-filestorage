package main

import (
	filestorage "github.com/jesseokeya/go-filestorage"
)

type Users struct {
	id        int
	firstName string
	lastName  string
	email     string
	password  string
	address   string
}

func main() {
	databaseName := "myAppUsers"
	schema := Users{
		id:        21334130,
		firstName: "Jesse",
		lastName:  "Okeya",
		password:  "encrypted",
		email:     "Jesseokeya@gmail.com",
		address:   "2550 Cotters Crescent, K1V8Y6",
	}

	dbInstance := filestorage.Store{
		databaseName,
		schema,
	}
}
