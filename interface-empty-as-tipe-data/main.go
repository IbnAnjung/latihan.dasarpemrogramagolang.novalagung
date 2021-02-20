package main

import "fmt"

func main() {
	// interface kosong akan menjadi tipe data yang bisa menyimpan nilai apa saja
	var secret interface{}
	var data map[string]interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	data = map[string]interface{}{
		"name":      "jhon wick",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}

	fmt.Println(data)
}
