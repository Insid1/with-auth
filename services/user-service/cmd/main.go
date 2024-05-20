package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// todo: move to env variables and read them
const (
	host     = "localhost"
	port     = 5440
	user     = "postgres"
	password = "postgres"
	dbname   = "go-auth-user"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	defer db.Close()
	if err != nil {
		log.Fatalf("Unable to open connection to DataBase. Error: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to connect to DataBase. Error: %s", err)
	}

	fmt.Println("Connected to DataBase")
}
