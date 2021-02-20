package main

import (
	"fmt"
	"os"
	"time"
)

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "second")
	os.Exit(0)
}

func main() {
	// fmt.Println("start")
	// time.Sleep(time.Second * 4)
	// fmt.Println("after 4  second")

	// // return objeck *time.Timer dengan property C bertime channel
	// var timer = time.NewTimer(4 * time.Second)
	// fmt.Println("start again")
	// <-timer.C
	// fmt.Println("end after 4 second more")

	// // time.AfterFunc
	// var ch = make(chan bool)
	// time.AfterFunc(4*time.Second, func() {
	// 	fmt.Println("expired")
	// 	ch <- true
	// })

	// fmt.Println("start again")
	// <-ch
	// fmt.Println("end again after 4 second")

	// // time.After
	// <-time.After(4 * time.Second)
	// fmt.Println("end after 4 second again")

	// done := make(chan bool)
	// ticker := time.NewTicker(2 * time.Second)

	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	done <- true
	// }()

	// for {
	// 	select {
	// 	case <-done:
	// 		ticker.Stop()
	// 		return
	// 	case t := <-ticker.C:
	// 		fmt.Println("Hello !!", t)
	// 	}
	// }

	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Print("what is 725/725 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right")
	} else {
		fmt.Println("the answer is wrong")
	}
}
