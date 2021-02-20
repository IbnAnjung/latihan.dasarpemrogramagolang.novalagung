package main

import (
	"fmt"
	"time"
)

func main() {

	var layoutFormat, value string
	var date time.Time

	var time1 = time.Now()
	fmt.Printf("time %v\n", time1)

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)

	var now = time.Now()
	fmt.Println("year :", now.Year(), "month :", now.Month())

	layoutFormat = "2006-01-02 15:04:05"
	value = "2006-09-02 14:13:12"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	layoutFormat = "02/01/2006 MST"
	value = "02/01/2020 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())
}
