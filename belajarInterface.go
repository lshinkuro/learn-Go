package main

import (
	"fmt"
)

type bangun2D interface {
	luas() float64
	keliling() float64
}

type bangun3D interface {
	volume() float64
}

type hitung interface {
	bangun2D
	bangun3D
}

type kubus struct {
	sisi float64
}

func (r *kubus) luas() float64 {
	return r.sisi * r.sisi
}

func (r *kubus) keliling() float64 {
	return r.sisi * 12
}

func (r *kubus) volume() float64 {
	return r.sisi * r.sisi * r.sisi
}

func main() {
	var K1 = kubus{4}
	var bangundatar hitung = &K1
	// bangundatar = &K1

	fmt.Println("luas kubus sama dengan ", bangundatar.luas())
	fmt.Println("Keliling kubus adalah", bangundatar.keliling())
	fmt.Println("volumenya kubus adalah", bangundatar.volume())

	fmt.Println("belajar embedded interface")
}
