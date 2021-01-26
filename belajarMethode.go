package main

import (
	"fmt"
)

type hitung interface {
	luas() float64

	keliling() float64
}

type persegipanjang struct {
	panjang, lebar float64
}

func (r persegipanjang) luas() float64 {
	return r.panjang * r.lebar
}

func (r persegipanjang) keliling() float64 {
	return r.panjang*2 + r.lebar*2
}

func main() {
	var P1 = persegipanjang{panjang: 8, lebar: 7}

	x := P1.luas()
	y := P1.keliling()

	fmt.Println(x, y)

	// membuat variabel interface

	var P2 hitung

	P2 = persegipanjang{12, 8}

	fmt.Println("luas persegi panjang adalah", P2.luas())
	fmt.Println("luas persegi panjang adalah", P2.keliling())

}
