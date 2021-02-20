package main

import (
	"fmt"
	"runtime"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(3)

	var message = make(chan string)

	/*
	* channel adalah pengirim ata penerima data
	* sifat dari channel adalah synchronous
	*
	 */
	// var sayHello = func(who string) {
	// 	var data = fmt.Sprintf("hello %s", who)
	// 	message <- data // synchronous
	// }

	// go sayHello("angga") // asynchronous
	// go sayHello("sadan") // asynchronous
	// go sayHello("Evie") // asynchronous

	// var message1 = <-message // synchronous
	// fmt.Println(message1)

	// var message2 = <-message // synchronous
	// fmt.Println(message2)

	// var message3 = <-message
	// fmt.Println(message3)

	// channel sebagai parameter
	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			message <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(message)
	}
}
