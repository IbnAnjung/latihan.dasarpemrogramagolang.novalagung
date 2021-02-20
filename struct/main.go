package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

type student struct {
	person //embed struct (nilai struct yang menjadi properti pada struct yang lain)
	age    int
	grade  int
}

// deklarasi struct 1 baris
//type teachear struct {person; study string}

// method pada struct
// (s student) adalah struct pemilik dari metod
func (s student) sayHello() {
	fmt.Println("Hallo, ", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func (s student) changeName(name string) {
	fmt.Println("----> on change name, name change to ", name)
	s.name = name
	fmt.Println("----> s.name after change in the method : ", s.name)
}

// pointer method
func (s *student) changeName2(name string) {
	fmt.Println("----> on change name2, name changed to ", name)
	s.name = name
}

func main() {

	// inisialisasi 1
	// var s1 student
	// s1.name = "jhon wick"
	// s1.grade = 2

	//inisialisasi 2
	//var s2 = student{"etahn", 2}

	// inisialisasi 3
	//var s3 = student{name: "jhonson"}

	//inisialisasi 4
	//var s4 = student{name: "wayne", grade: 3}

	//inisialisasi 5
	//var s5 = student{grade: 5, name: "bruce"}

	// pointer
	//var s6 *student = &s1

	// fmt.Println("s1 name :", s1.name)
	// fmt.Println("s2 name :", s2.name)
	// fmt.Println("s3 name :", s3.name)
	// fmt.Println("s4 name :", s4.name)
	// fmt.Println("s5 name :", s5.name)
	// fmt.Println("s6 name :", s6.name)

	// s6.name = "wick jhon"

	// fmt.Println("s6 name :", s6.name)
	// fmt.Println("s1 name :", s1.name)

	// embend struct
	var s1 = student{}
	s1.name = "Jhon wick"
	s1.age = 25        // age of student
	s1.person.age = 20 // age of person
	s1.grade = 1

	var p2 = person{name: "jhon", age: 19}
	var s2 = student{person: p2, grade: 2}

	// anonymous struct
	var srt1 = struct {
		person
		grade int
	}{}
	srt1.person = person{"anonymous", 21}
	srt1.grade = 1

	var srt2 struct {
		person
		grade int
	}

	srt2.person = person{"anonymous 2", 22}
	srt2.grade = 2

	var srt3 = struct {
		grade int
	}{
		grade: 12,
	}

	//slice of struct
	var allPerson = []person{
		{name: "slice 1", age: 1},
		{name: "slice 2", age: 2},
		{name: "slice 3", age: 3},
	}

	// slice od anonymous struct
	var allStudent = []struct {
		person
		grade int
	}{
		{person: person{name: "student 1", age: 1}, grade: 1},
		{person: person{name: "student 2", age: 2}, grade: 2},
		{person: person{name: "student 3", age: 3}, grade: 3},
	}

	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.person.age)
	fmt.Println(s1.grade)

	fmt.Println(s2.name)
	fmt.Println(s2.age)
	fmt.Println(s2.person.age)
	fmt.Println(s2.grade)

	fmt.Println(srt1.name)
	fmt.Println(srt1.age)
	fmt.Println(srt1.person.age)
	fmt.Println(srt1.grade)

	fmt.Println(srt2.name)
	fmt.Println(srt2.age)
	fmt.Println(srt2.person.age)
	fmt.Println(srt2.grade)

	fmt.Println(srt3.grade)

	for _, pers := range allPerson {
		fmt.Println(pers.name, " age is ", pers.age)
	}

	for _, studt := range allStudent {
		fmt.Println(studt.person.name, " age is ", studt.person.age, " grade is ", studt.grade)
	}

	// penggunaan method
	s1.sayHello()
	var name = s1.getNameAt(2)
	fmt.Println(name)

	s1.changeName("wick Jhon")
	fmt.Println("----> s.name after change out side the method : ", s1.name)

	s1.changeName2("wick Jhon")
	fmt.Println("----> s.name after change2 out side the method : ", s1.name)

}
