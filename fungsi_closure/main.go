package main

import "fmt"

func main() {
	// example 1
	var numbers = []int{2, 3, 4, 3, 4, 2, 3}

	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max)

	//example 2
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}

			r = append(r, e)
		}

		return r
	}(4)
	fmt.Println("Original number: ", numbers)
	fmt.Println("Filtered numbers: ", newNumbers)

	//example 3, closure sebagai callback
	var maxNumber = 3
	var howMany, getNumbers = findMax(numbers, maxNumber)
	var theNumbers = getNumbers()
	fmt.Println("numbers \t", numbers)
	fmt.Printf("find \t: %d\n\n", maxNumber)
	fmt.Println("found \t:", howMany)
	fmt.Println("value \t:", theNumbers)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	/* contoh 1
	return len(res), func() []int{
		return res
	}
	*/
	// bisa juga di masukan ke variable dulu
	var getNumbers = func() []int {
		return res
	}

	return len(res), getNumbers
}
