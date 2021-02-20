package main

import "fmt"

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func main() {
	//var avg = calculate(2, 3, 1, 2, 3, 4, 5, 6, 7, 7, 7, 7, 4)
	var numbers = []int{2, 3, 1, 2, 3, 4, 5, 6, 7, 7, 7, 7, 4}
	var avg = calculate(numbers...)
	var msg = fmt.Sprintf("Rata - Rata: %.2f", avg)
	fmt.Println(msg)
}
