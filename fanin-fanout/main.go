package main

import (
	"fanin-fanout/sequencial"
	"log"
	"time"
)

func main() {
	log.Println("start")
	start := time.Now()

	sequencial.GenerateFile()
	// concuren.GenerateFile()
	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}
