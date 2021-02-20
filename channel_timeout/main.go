package main

import (
	"fmt"
	"math/rand"
	"time"
	"runtime"
)

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retreiveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("timeout. no activities under 5 second")
			break loop
		}
	}
}

func main(){
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)
	
	var message = make(chan int)

	go sendData(message)
	retreiveData(message)
}