package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "jhon wick"

	// var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	// fmt.Println("encoded:", encodedString) //encoded: amhvbiB3aWNr

	// var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	// var decodedString = string(decodedByte)
	// fmt.Println("decoded:", decodedString) //decoded

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString = string(encoded)
	fmt.Println(encodedString)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString = string(decoded)
	fmt.Println(decodedString)

	var url = "https://kalipare.com"
	var encodedUrl = base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println(encodedUrl)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedUrl)
	var decodedUrl = string(decodedByte)
	fmt.Println(decodedUrl)

}
