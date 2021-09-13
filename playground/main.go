package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var wg *sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(3)
	start := time.Now()

	wg = new(sync.WaitGroup)
	content := write()

	ls1 := listen(content, "ls1")
	ls2 := listen(content, "ls2")
	ls3 := listen(content, "ls3")
	ls4 := listen(content, "ls4")
	listenMerge(ls1, ls2, ls3, ls4)
	wg.Wait()

	fmt.Println("finish ", time.Since(start).Seconds())
}

func listenMerge(chanInMany ...<-chan string) {
	wg.Add(len(chanInMany))
	for _, each := range chanInMany {
		go func(each <-chan string) {
			log.Println("start listen channel")
			for str := range each {
				fmt.Printf("add string %s\n", str)
			}
			log.Println("finish listen channel")
			wg.Done()
		}(each)
	}
}

func write() <-chan string {
	chanOut := make(chan string)

	go func() {
		log.Println("start writing")
		for i := 1; i <= 20; i++ {
			chanOut <- strconv.Itoa(i)

		}
		log.Println("finish writing")
		close(chanOut)
	}()

	return chanOut
}

func listen(chanIn <-chan string, chs string) <-chan string {
	chanOut := make(chan string)
	go func() {
		log.Println("start listen ch", chs)
		for str := range chanIn {
			chanOut <- fmt.Sprintf("from listen %s-%s", str, chs)
		}
		log.Println("finish listen ch", chs)
		close(chanOut)
	}()

	return chanOut
}
