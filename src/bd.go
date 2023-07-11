package app

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type databaseType struct {
	db *sql.DB
	tx *sql.Tx
}

type dataError

func (database *databaseType) openConnection() error {

	var err error
	var user, password string
	err = getUserandPassword(&user, &password)
	if(err != nil){

	}
	database.db, err = sql.Open("postgres", "host=localhost port=5432 user=palma password=123456 dbname=teste sslmode=disable")
	if err != nil {
		fmt.Println("Error opening database!")
		panic(err)
	}

	return err
}

func getUserandPassword(user *string, password *string) int {

	var err error

	buffer, err := os.ReadFile("../assets/credentials.dat")

	return ERROR_READING_LOCAL_FILE
}
