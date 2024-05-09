package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        age INT NOT NULL
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertUser(db *sql.DB, name string, age int) (int, error) {
	sqlStatement := `
    INSERT INTO users (name, age)
    VALUES ($1, $2)
    RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, name, age).Scan(&id)
	if err != nil {
		return 0, err
	}
	fmt.Println("New record ID is:", id)
	return id, nil
}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=postgres password=mysecretpassword sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable(db)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			var user User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := insertUser(db, user.Name, user.Age)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			user.ID = id
			response, err := json.Marshal(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)

		case "GET":
			users, err := getUsers(db)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			response, err := json.Marshal(users)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	})

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
