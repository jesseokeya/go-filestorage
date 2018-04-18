package main

import (
	"fmt"

	filestorage "github.com/jesseokeya/go-filestorage"
)

// schema struct is a schema for app users
type schema struct {
	ID        int    `json:"_id,omitempty"`
	FirstName string `json:"firtsName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Address   string `json:"address,omitempty"`
}

func main() {
	// user := schema{
	// 	ID:        21334131,
	// 	FirstName: "Jesse",
	// 	LastName:  "Okeya",
	// 	Password:  "encrypted",
	// 	Email:     "Jesseokeya@gmail.com",
	// 	Address:   "2550 Cotters Crescent, K1V8Y6",
	// }

	database := filestorage.Connect()
	database.Name("Jesse's Users")

	fmt.Println(database)
}
