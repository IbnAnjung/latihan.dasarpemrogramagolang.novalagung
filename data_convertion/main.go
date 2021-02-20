package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "124"
	var nu = 123

	// string to numeric
	var num, err = strconv.Atoi(str)
	if err == nil {
		fmt.Println(num)
	}

	// numeric to str
	var st = strconv.Itoa(nu)
	fmt.Println(st)

	// 10 basis angka (decimal)
	// 64 panjang int(int64)
	_, _ = strconv.ParseInt(str, 10, 64)
	// 2 basis angka (biner)
	// 8 panjang int(int 8)
	_, _ = strconv.ParseInt(str, 2, 8)

	// string ke float
	_, _ = strconv.ParseFloat("24.12", 32) // 32 lebar decimal

	// convert int ke string dengan basis numerik bisa di tentukan
	str = strconv.FormatInt(int64(24), 8)  // basis 8
	str = strconv.FormatInt(int64(24), 2)  // basis 2 biner
	str = strconv.FormatInt(int64(24), 10) // basis 10 decimal
	// float ke string
	str = strconv.FormatFloat(float64(24.12), 'f', 6, 64) // lebar decimal 6, basis float64

	// string ke bool
	var bul, _ = strconv.ParseBool("true")
	fmt.Println(bul)
	// bool ke string
	strconv.FormatBool(true)

	// casting string <-> byte
	var text1 = "hallo"
	var b = []byte(text1)

	fmt.Printf("%d %d %d %d %d", b[0], b[1], b[2], b[3], b[4])

	
}
