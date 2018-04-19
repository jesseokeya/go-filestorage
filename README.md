# go-filestorage
## An in memory file storage system implemented in go. 

## Installation
`go get -u github.com/jesseokeya/go-filestorage`

```go
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
	users := []schema{
		{ID: 064245,
			FirstName: "Peter",
			LastName:  "Smith",
			Password:  "encrypted",
			Email:     "peter_smith@gmail.com",
			Address:   "2550 New Brick Avenue, Washington Dc",
		},
		{ID: 167215,
			FirstName: "John",
			LastName:  "Mendez",
			Password:  "encrypted",
			Email:     "johnmendez@yahoo.com",
			Address:   "344 Smithinson Road, Manitoba",
		},
		{ID: 31290,
			FirstName: "Wang",
			LastName:  "Lee",
			Password:  "encrypted",
			Email:     "wang_lee@gmail.com",
			Address:   "440 Handerson Avenue, Houston",
		},
	}

	database := filestorage.Connect()
	database.Name("App Users")
	fmt.Println(database.GetPath())
	for _, item := range users {
		database.InsertOne(item)
	}
	fmt.Println(database.FindAll())
}
```

## Editor Snippet
<img src="https://github.com/jesseokeya/go-filestorage/blob/master/image/editor.png" width="500" height="500">
