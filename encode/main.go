package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "jhon wick"

	// simple encode
	// var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	// fmt.Println("encoded:", encodedString)

	// var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	// var decodedString = string(decodedByte)
	// fmt.Println("decoded byte:", decodedByte)
	// fmt.Println("decoded:", decodedString)

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString = string(encoded)
	fmt.Println(encodedString)
}
