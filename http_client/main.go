package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUser(ID string)(student, error){
	var err error
	var client = &http.Client{}
	var data student

	// var param = url.Values{}
	// param.Set("id", ID)
	// var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("post", baseURL+"/user", payload)
	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil{
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil{
		return data, err
	}

	return data, nil
}
func main() {
	var users, err = fetchUsers()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	var user,_  = fetchUser("E001")
	

	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}
	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user.ID, user.Name, user.Grade)
}
