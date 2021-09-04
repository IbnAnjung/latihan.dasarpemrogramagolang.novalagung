package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "Jhon Wick",
			"Message": "have a nice day",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/index", index)
	fmt.Println("starting web server at http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
