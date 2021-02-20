package main

import (
	f "fmt"
	"level-akses/library"
)

func main() {
	library.SayHello()
	// contoh pemanggilan method un-exported
	// it will be return error refer to unecported name librarytou
	//library.introduce("ethan")
	library.SayHelloAndIntroduce("ethan")

	// level akses pada struct
	var s1 = library.Student{"ethan", 21}
	f.Println("name", s1.Name)
	f.Println("grade", s1.Grade)

	f.Printf("Name : %s\n", library.Std.Name)
	f.Printf("Grade : %d\n", library.Std.Grade)
}
