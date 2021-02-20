package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "Jhon Wick", "Age": 27}`
	var arrayJsonString = `[
		{"Name":"Jhon Wick", "Age":27s}
	]`
	var jsonData = []byte(jsonString)
	var data1 map[string]interface{}
	var data2 interface{}
	var data User

	var err = json.Unmarshal(jsonData, &data)
	json.Unmarshal(jsonData, &data1)
	json.Unmarshal(jsonData, &data2)
	var decodeData2 = data2.(map[string]interface{})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User:", data.FullName)
	fmt.Println("age:", data.Age)
	fmt.Println("user:", data1["Name"])
	fmt.Println("age:", data1["Age"])
	fmt.Println("user:", decodeData2["Name"])
	fmt.Println("age:", decodeData2["Age"])

}
