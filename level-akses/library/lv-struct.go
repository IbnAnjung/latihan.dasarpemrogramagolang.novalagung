package library

import "fmt"

// // struct unexported
// type student struct {
// 	Name  string
// 	grade int
// }

// struct exported
type Student struct {
	Name  string
	Grade int
}

var Std = Student{}

func init() {
	Std.Name = "Jhon Wick"
	Std.Grade = 2
	fmt.Println("----> library/lv-struct.go imported")
}
