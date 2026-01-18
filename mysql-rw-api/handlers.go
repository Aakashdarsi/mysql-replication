package main

import (
	"encoding/json"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := writeDB.Exec(
			"INSERT INTO users(id, name) VALUES (?, ?)",
			u.ID, u.Name,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("inserted"))

	case http.MethodGet:
		rows, err := readDB.Query("SELECT id, name FROM users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []User

		for rows.Next() {
			var u User
			rows.Scan(&u.ID, &u.Name)
			users = append(users, u)
		}

		json.NewEncoder(w).Encode(users)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
