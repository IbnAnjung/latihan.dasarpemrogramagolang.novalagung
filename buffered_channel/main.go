// buffered channel sama dengan channel dengan nilai n (banyaknya channel) untuk proses asynchronous

package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	message := make(chan int, 2)
	// 2 adalah nilai buffer / banyaknya channel untuk proses asynchronous

	go func() {
		for {
			i := <-message
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		message <- i
	}
}
