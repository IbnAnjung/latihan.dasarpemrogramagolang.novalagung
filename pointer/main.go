package main

import "fmt"

func main() {
	var numberA = 4
	/*
		*int => bertipe pointer dengan value int
		&numberA => referencing to variable numberA (mengambil nilai pointer variable numberA)
	*/
	var numberB *int = &numberA

	fmt.Println("numberA       = ", numberA)
	fmt.Println("numberA (add) = ", &numberA)
	/*
	 *numberB => derefencing numberB (mengambil nilai value dari numberB)

	 */
	fmt.Println("numberB       = ", *numberB)
	fmt.Println("numberB (add) = ", numberB)

	change(&numberA, 10)
	fmt.Println("numberB after change = ", numberA)
	fmt.Println("numberB after change = ", *numberB)
}

func change(original *int, value int) {
	*original = value
}

/*
 * variable numberA bernilai 4
 * variable numberB adalan pointer dengan referensi alamat dari varible numberB
 * hal ini membuat nilai asli dari numberB akan terus berubah ketika niali variable numberA berubah
 */
