package main

import (
	"database/sql"
	"fmt"
	"os"
	"errors"

	_ "github.com/lib/pq"
)

type databaseType struct {
	db *sql.DB
	tx *sql.Tx
	tableName string
}

type employee struct {
	pk            int
	nome          string
	salario       int
	cd_nascimento string
	idade         int
}

type accounts {
	account_id      int
	nome          	string
	cd_agencia 		string
	saldo        	float64
	status 		  	string
	cliente_since 	string
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

func (database *databaseType) createTable() sql.Result {

	result := exec(database.db, "drop table if exists "+ database.tableName)

	if result != nil {
		fmt.Println("\nCreating table!\n")

		result = exec(database.db, `create table `+ database.tableName+ ` (
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

func (database *databaseType) insert(e *accounts) error {

	exist,err := checkAccount(e)

	if exist == true{
		fmt.Println(err)
		return errors.New("This account already existe !!")
	}

	stmt, err := database.tx.Prepare("insert into accounts (id_account, nome , cd_agencia, saldo, status, cliente_since) values ($1, $2, $3, $4, $5, $6)")

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

func (database *databaseType) checkAccount(e *accounts) bool, error{

	buffer := selectSpecific("account_id, name", "accounts")
	var aux accounts

	for buffer.Next() {
		rows.Scan(&aux.account_id, &aux.nome)
		if(aux.account_id == e.account_id){
			return true, errors.New("Account " + aux.nome + "already exist!")
		}
	}
}

func (database *databaseType) selectSpecific(statement string, tableName string) {

	rows, _ := database.db.Query("select * from " + tableName)
	defer rows.Close()

	for rows.Next() {
		var e employee
		rows.Scan(&e.pk, &e.nome, &e.salario, &e.cd_nascimento, &e.idade)
		fmt.Println(e)
	}
}


func (database *databaseType) selectAll(tableName string) {

	rows, _ := database.db.Query("select * from " + tableName)
	defer rows.Close()

	for rows.Next() {
		var e employee
		rows.Scan(&e.pk, &e.nome, &e.salario, &e.cd_nascimento, &e.idade)
		fmt.Println(e)
	}
}
