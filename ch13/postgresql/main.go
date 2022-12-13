package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "datran"
	password = "danghai42Paris"
	dbname   = "golang"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect database successfully.")
	employeeTable := `
	CREATE TABLE IF NOT EXISTS employees2022 (
		PersonID int,
		LastName varchar(255),
		FirstName varchar(255),
		Address varchar(255),
		City varchar(255)
	);`
	sql, err := db.Exec(employeeTable)
	if err != nil {
		panic(err)
	}
	fmt.Println("The table is created")
	fmt.Println(sql)
	insertRow1 := `
	INSERT INTO employees2022 (PersonID, LastName, FirstName, Address, City)
	VALUES (1, 'Henry', 'M.', '121 - 14th Ave.S.Suite 3B', 'Seattle');`
	sql, err = db.Exec(insertRow1)
	if err != nil {
		panic(err)
	}
	fmt.Println("1st Row is inserted.")
	fmt.Println(sql)
}
