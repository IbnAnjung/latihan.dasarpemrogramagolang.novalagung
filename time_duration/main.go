package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	time.Sleep(5 * time.Second)
	end := time.Now()

	duration := time.Since(start)
	duration2 := start.Sub(end)

	// tipe varibale harus time.Duration untuk mengalikan jika ditampung ke variable
	var n time.Duration = 5
	dur := n * time.Second
	fmt.Orintln(dur)
	// konversi angka ke time.Duration
	/*
	* 12 minute = 12 * time.Minute
	* 65 jam = 65 * time.Hour
	* 150k minute = 150000 * time.Milisecond
	* 45 micro detik = 45 * time.Microsecond
	* 233 nano detik = 233 * Nanosecond
	 */

	fmt.Println("time elapsed in second", duration.Seconds())
	fmt.Println("time elapsed in minute", duration.Minutes())
	fmt.Println("time elapsed in hours ", duration.Hours())

	fmt.Println("time elapsed in second", duration2.Seconds())
	fmt.Println("time elapsed in minute", duration2.Minutes())
	fmt.Println("time elapsed in hours ", duration2.Hours())
}
