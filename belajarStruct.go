package main

import (
	"fmt"
)

type person struct {
	name, address string
	age           int
	hobbies       []string
}

type student struct {
	grade int
	person
}

func (r *person) ChangeName(name string) {
	r.name = name
}

func (r student) ShowIdentity() string {
	return r.name + r.address
}

func main() {
	var P1 = person{}

	P1.name = "agung"
	P1.address = "bandung"
	P1.age = 34
	P1.hobbies = []string{"lari", "berenang", "ngudud"}

	fmt.Println("hari ini kita belajar tentang struct")

	fmt.Println("nama saya adalah :", P1.name)
	fmt.Println("umur saya adalah :", P1.age)
	fmt.Println("alamat saya adalah :", P1.address)

	fmt.Println("hobby saya adalah")

	for i, v := range P1.hobbies {
		fmt.Println(i+1, v)
	}

	P1.ChangeName("ahmad")

	fmt.Println(P1.name)

	P2 := person{name: "bagas", address: "kalimantan", age: 21}

	fmt.Println(P2.age, P2.name, P2.address)

	var P3 *person = &P2

	P3.age = 78

	fmt.Println(P3.age, P3.name, P3.address)

	var S1 = student{person: P2, grade: 20}

	var x = S1.ShowIdentity()
	fmt.Println(x)

	fmt.Println(S1.age)

}
