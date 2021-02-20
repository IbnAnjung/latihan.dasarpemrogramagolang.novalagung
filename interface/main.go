package main

import (
	"fmt"
	"math"
)

// normally interface
// type hitung interface {
// 	luas() float64
// 	keliling() float64
// }

// embed interface

type hitung2D interface {
	luas() float64
	keliling() float64
}

type hitung3D interface {
	volume() float64
}

type hitung interface {
	hitung2D
	hitung3D
	greeting() string
}

type bola struct {
	diameter float64
}

func (l bola) jariJari() float64 {
	return l.diameter / 2
}

func (l bola) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l bola) keliling() float64 {
	return math.Pi * l.diameter
}

func (l bola) volume() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l bola) greeing() string {
	return "hai aku bola"
}

type kubus struct {
	sisi float64
}

func (p kubus) luas() float64 {
	return math.Pow(p.sisi, 2) * 6
}

func (p kubus) keliling() float64 {
	return p.sisi * 4 * 6
}

func (p kubus) volume() float64 {
	return math.Pow(p.sisi, 3)
}

func (p kubus) greeting() string {
	return "Hai aku kubus"
}

func main() {
	var bangunDatar hitung

	bangunDatar = kubus{sisi: 10.0}
	fmt.Println("=============== persegi")
	fmt.Println(bangunDatar.greeting())
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("keliling  :", bangunDatar.volume())
}
