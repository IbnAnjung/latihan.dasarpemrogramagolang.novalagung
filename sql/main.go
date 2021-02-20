package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:CARDO@tcp(127.0.0.1:3306)/db_belajar_golang")
	// with postgre sql
	// db, err := sql.Open("pq", "root:CARDO@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
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

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return
	}

	for _, each := range result {
		fmt.Println(each.id, ".", each.name, each.grade)
	}
}

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student{}
	var id = "E001"

	err = db.
		QueryRow("select name, grade from tb_student where id = ?", id).
		Scan(&result.name, &result.grade)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name: %s \ngrade: %d\n", result.name, result.grade)
}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	var result2 = student{}

	stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	stmt.QueryRow("B002").Scan(&result2.name, &result2.grade)

	fmt.Println("_____ this result form sql prepare ________")
	fmt.Printf("name: %s\ngrade: %d\n", result1.name, result1.grade)
	fmt.Printf("name: %s\ngrdae: %d\n", result2.name, result2.grade)

}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.
		Exec("insert into tb_student values(?, ?, ?, ?)", "G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success")

	_, err = db.
		Exec("update tb_student set age =  ? where id = ?", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success")

	// dengan prepare statement
	stmt, err := db.Prepare("delete from tb_student where id=?")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	stmt.Exec("G001")
	fmt.Println("delete success, with prepare statement")
}

func main() {
	sqlQueryRow()
	sqlQuery()
	sqlPrepare()
	sqlExec()
}
