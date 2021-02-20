package main

import (
	"fmt"
	"runtime"
)

// parameter ch hanya untuk mengirim data
func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

// parameter ch hanya untuk menerima data
func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)
	go sendMessage(message)
	printMessage(message)
}
