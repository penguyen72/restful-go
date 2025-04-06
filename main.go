package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "myuser"
    password = "mypassword"
    dbname   = "mydb"
)

func signInHandler(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
							host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	pingError := db.Ping()
	if pingError != nil {
		log.Fatal("Error Connecting:", err)
	}

	fmt.Println("Connected!")
}

func signOutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}

func main() {
	http.HandleFunc("/sign-in", signInHandler)
	http.HandleFunc("/sign-out", signOutHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/get-users", getUsersHandler)

	http.ListenAndServe(":8080", nil)
}