package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error Occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}

	return true, nil
}

func main() {

	defer catch()

	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}

	if valid, err := validate(name); valid {
		fmt.Println("Hallo", name)
	} else {
		//fmt.Println(err.Error())
		//panic errro (semua proses stelahnya akan di hentikan kecuali yang di defer)
		panic(err.Error())
		fmt.Println("ini tidak akan di tampilkan")
	}
}
