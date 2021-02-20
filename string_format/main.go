package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

func main() {
	var data = student{
		name:        "angga",
		height:      65.5,
		age:         26,
		isGraduated: true,
		hobbies:     []string{"programming", "travelling"},
	}

	// %b => numerik biner
	fmt.Printf("%b\n", data.age)

	// %x / %X => numerik heksa desimal
	fmt.Printf("%x\n", data.age)

	// %d => number 10 basis
	fmt.Printf("%d\n", data.age)

	// %c => unicode menjadi kode
	fmt.Printf("%c\n", 1400)

	// $e / %e numerik decimal ke standart Scientific notation
	fmt.Printf("%e\n", data.height)

	// %f / $F numerik decimal => defaultnya 6 decimal
	fmt.Printf("%f\n", data.height)
	fmt.Printf("%.2f\n", data.height)
	fmt.Printf("%.f\n", data.height)

	// %g / %G => numerik decimal, cocok jika decimlanya sangat panjang, sesuai dengan data
	fmt.Printf("%g\n", 12.32143244546576876879)

	// %0 => numerik octal
	fmt.Printf("%o\n", data.age)

	// %p => alamat pointer
	fmt.Printf("%p\n", data.age)

	// %s => data string
	fmt.Printf("%s\n", data.name)

	// %t => boolean
	fmt.Printf("%t\n", data.isGraduated)

	// $T => mengembalikan string tipe variable yang di format
	fmt.Printf("%T\n", data.name)
	fmt.Printf("%T\n", data.age)

	// %v	=> mengembalikan nilai string tipe data apa saja
	fmt.Printf("%v\n", data)

	// %+v => mengembalikan format struct untuk tiap property dan nilainya
	fmt.Printf("%+v\n", data)

	// %#v sama seperti %+v + bagaimana struct di deklarasikan
	fmt.Printf("%#v\n", data)

}
