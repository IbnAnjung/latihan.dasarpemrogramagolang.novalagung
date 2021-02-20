package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// rand.Seed(10)
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("random ke-1 :", rand.Int())
	fmt.Println("random ke-2 :", rand.Int())
	fmt.Println("random ke-3 :", rand.Int())

	// random string
	var letters = []rune("acdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYX")

	var randomString = func(length int) string {
		b := make([]rune, length)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}

	fmt.Println(randomString(10))
	fmt.Println(rand.Intn(10))
}
