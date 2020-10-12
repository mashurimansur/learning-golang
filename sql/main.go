package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //nolint:unused
)

type Student struct {
	ID    string
	Name  string
	Age   int
	Grade int
}

func connect() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}
	return
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Student

	for rows.Next() {
		var each = Student{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.Name)
	}
}

func main() {
	sqlQuery()
}
