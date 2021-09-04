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

var DB *sql.DB
var Err error
var Student student

func connect() {
	DB, Err = sql.Open("mysql", "angga:angga@tcp(127.0.0.1:3306)/learn_golang")
	if Err != nil {
		fmt.Println(Err.Error())
		return
	}

}

func insertWithExec(student student) {
	connect()
	defer DB.Close()

	fmt.Print("Try to insert new row => ")

	_, Err := DB.Exec("INSERT INTO tb_student VALUES (?, ?, ?, ?)", student.id, student.name, student.age, student.grade)
	if Err != nil {
		fmt.Println(Err.Error())
		return
	}

	fmt.Println("success")
}

func sqlQuery() {
	connect()
	defer DB.Close()
	fmt.Println("Try to insert new row")

	var age = 27
	rows, Err := DB.Query("select id, name, grade from tb_student where age = ?", age)
	if Err != nil {
		fmt.Println(Err.Error())
		return
	}
	defer rows.Close()

	var result []student
	for rows.Next() {
		var each = student{}
		Err = rows.Scan(&each.id, &each.name, &each.grade)
		if Err != nil {
			fmt.Println(Err.Error())
			return
		}
		result = append(result, each)
	}

	if Err = rows.Err(); Err != nil {
		fmt.Println(Err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

func updateWithPrepare(student student) {
	connect()
	defer DB.Close()

	fmt.Print("try to update with prepare statement => ")

	stmt, Err := DB.Prepare("UPDATE tb_student SET name = ?, age = ?, grade = ? WHERE id = ?")
	if Err != nil {
		fmt.Println(Err.Error())
		return
	}
	stmt.Exec(student.name, student.age, student.grade, student.id)
	fmt.Println("Success")
}

func delete(id string) {
	connect()
	defer DB.Close()

	fmt.Println("try to delete row")

	stmt, Err := DB.Prepare("DELETE FROM tb_student where id = ?")
	if Err != nil {
		fmt.Println(Err.Error())
		return
	}
	stmt.Exec(id)
}

func sqlPrepare(id string) {
	connect()
	defer DB.Close()

	fmt.Println("Try to get data")
	stmt, err := DB.Prepare("select id, name, grade, age from tb_student where id = ?")
	if err != nil {
		fmt.Println(Err.Error())
		return
	}

	stmt.QueryRow(id).Scan(&Student.id, &Student.name, &Student.age, &Student.grade)
}

func main() {
	Student = student{
		id:    "B006",
		name:  "GALAHAD",
		age:   30,
		grade: 12,
	}
	insertWithExec(Student)
	sqlPrepare(Student.id)
	fmt.Println("===== new Rows =====")
	fmt.Printf("ID: %s\nName: %s\nAge: %d\nGrade: %d\n", Student.id, Student.name, Student.age, Student.grade)
	fmt.Println("====================")
	Student.name = "Shonyo"
	Student.age = 25
	updateWithPrepare(Student)
	sqlPrepare(Student.id)
	fmt.Println("===== updated Rows =====")
	fmt.Printf("ID: %s\nName: %s\nAge: %d\nGrade: %d\n", Student.id, Student.name, Student.age, Student.grade)
	fmt.Println("====================")
	delete(Student.id)

}
