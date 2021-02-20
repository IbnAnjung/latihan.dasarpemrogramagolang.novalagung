package main

import (
	"fmt"
	"strings"
)

// alias closure function
type FilterCallback func(string) bool

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}

// filter with alias
func filterWithAlias(data []string, callback FilterCallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}

func main() {
	var data = []string{"wick", "jason", "ethan", "indhi"}

	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	var dataContainsOAlias = filterWithAlias(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLength5Alias = filterWithAlias(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli: \t\t", data)
	fmt.Println("filter yang mengangun huruf \"o\" \t:", dataContainsO)
	fmt.Println("filter data dengan panjang \"5\" \t:", dataLength5)

	fmt.Println("filter with alias yang mengangun huruf \"o\" \t:", dataContainsOAlias)
	fmt.Println("filter with alias data dengan panjang \"5\" \t:", dataLength5Alias)
}
