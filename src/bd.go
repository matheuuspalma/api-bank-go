package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type databaseType struct {
	db *sql.DB
	tx *sql.Tx
}

type employee struct {
	pk            int
	nome          string
	salario       int
	cd_nascimento string
	idade         int
}

func (database *databaseType) openConnection() int {
	var user, password string

	err := getUserandPassword(&user, &password)

	CheckError(err)

	database.db, err = sql.Open("postgres", "host=localhost port=5432 user="+user+" password="+password+"dbname=teste sslmode=disable")

	CheckError(err)

	return 0
}

func main() {
	fmt.Println("Testing this code!!")

	var user, senha string

	getUserandPassword(&user, &senha)

	fmt.Println("User = " + user + "Password = " + senha)
}

func getUserandPassword(user *string, password *string) error {

	var err error
	var isToCopy bool = false
	var userWereCopy bool = false

	bufferByte, err := os.ReadFile("../assets/credentials.dat")

	if err == nil {
		buffer := string(bufferByte)
		for _, c := range buffer {
			if c == '=' {
				isToCopy = true
			} else if c == '\n' {
				isToCopy = false
				userWereCopy = true
			} else if isToCopy == true {
				if userWereCopy == true {
					*password += string(c)
				} else {
					*user += string(c)
				}
			}
		}
	} else {
		fmt.Println("File not founded :( !")
	}
	return err
}

func createTable(db *sql.DB, table string) sql.Result {

	result := exec(db, "drop table if exists "+table)

	if result != nil {
		fmt.Println("\nCreating table!\n")

		result = exec(db, `create table `+table+` (
		primary_key int,
		nome varchar(80),
		salario int,
		cd_nascimento varchar(30),
		idade int
	)`)
	}

	if result == nil {
		fmt.Println("Error executing task on database.")
	}
	return result
}

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func insert(tx *sql.Tx, e *employee) error {

	stmt, err := tx.Prepare("insert into employees(primary_key, nome, salario, cd_nascimento, idade) values ($1, $2, $3, $4, $5)")

	if err != nil {
		fmt.Printf("Error inserting employee!")
		return err
	}
	fmt.Println("Inserindo employee " + e.nome)

	res, err := stmt.Exec(e.pk, e.nome, e.salario, e.cd_nascimento, e.idade)

	linhas, _ := res.RowsAffected()
	fmt.Printf("Inserted : %d \n", linhas)

	return err
}

func selectAll(db *sql.DB, tableName string) {

	rows, _ := db.Query("select * from " + tableName)
	defer rows.Close()

	for rows.Next() {
		var e employee
		rows.Scan(&e.pk, &e.nome, &e.salario, &e.cd_nascimento, &e.idade)
		fmt.Println(e)
	}
}
