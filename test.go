package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	jsonString := `{"title":"Learning JSON in Golang","author":"Lanka"}`
	var book Book
	err := json.Unmarshal([]byte(jsonString), &book)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", book)
}
