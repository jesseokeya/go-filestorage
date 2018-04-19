package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	filestorage "github.com/jesseokeya/go-filestorage"
	httplogger "github.com/jesseokeya/go-httplogger"
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

var (
	database = filestorage.Connect()
)

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
	database.Name("App Users")
	fmt.Println(database.GetPath())
	fmt.Println(database.GetName())
	for _, item := range users {
		database.InsertOne(item)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	fmt.Println("Server running on port *" + "8000")
	http.ListenAndServe(":8000", httplogger.Golog(r))
}

// HomeHandler handle home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(database.FindAll())
	if err != nil {
		panic(err)
	}
	w.Write(data)
}
