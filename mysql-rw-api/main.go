package main

import (
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	initDB()

	http.HandleFunc("/users", UsersHandler)

	log.Println("API running on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
