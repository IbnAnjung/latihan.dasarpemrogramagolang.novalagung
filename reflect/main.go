package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variable :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable :", reflectValue.Int())
		fmt.Println("tipe variable :", reflectValue.Type())
		fmt.Println("nilai variable :", reflectValue.Interface())
	}

	var s1 = student{Name: "angga saputra", Grade: 1}
	s1.getPropertyInfo()

	var s2 = &stdent{Name: "jhon wick", Grade: 2}
	fmt.Println("nama             :", s2.Name)

	var reflectValue = relect
}
