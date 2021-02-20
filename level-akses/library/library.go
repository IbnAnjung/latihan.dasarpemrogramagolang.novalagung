package library

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}

func SayHelloAndIntroduce(name string) {
	fmt.Println("Hello")
	introduce(name)
}

func introduce(name string) {
	fmt.Println("nama saya", name)
}
