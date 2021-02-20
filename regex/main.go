package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "banana burger soup"

	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)

	// cek kecocokan string dengan pola regex
	isMatch := regex.MatchString(text)
	fmt.Println(isMatch)

	stringMatch := regex.FindString(text)
	fmt.Println(stringMatch)

	stringMatchIndex := regex.FindStringIndex(text)
	fmt.Println(stringMatchIndex)

	repalceStringMatch := regex.ReplaceAllString(text, "match")
	fmt.Println(repalceStringMatch)

	replaceStringMatchWithFunc := regex.ReplaceAllStringFunc(text, func(each string) string {
		if each != "burger" {
			return "match"
		}

		return each
	})
	fmt.Println(replaceStringMatchWithFunc)

	regex, err = regexp.Compile(`[a-b]+`)
	split := regex.Split(text, -1)
	fmt.Printf("%#v \n", split)
}
