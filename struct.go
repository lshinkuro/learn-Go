package main

import (
	"fmt"
)

// cara membuat struct 1
type Address struct {
	id           int
	nama, alamat string
}

func (u Address) Name() string {
	return u.nama
}

func (u *Address) setName(newName, newAlamat string) {
	u.nama = newName
	u.alamat = newAlamat
}

// cara membuat struct 2
type Numb struct {
	id, NoTogel int
}

func (u Numb) Num() int {
	return u.id
}

func (u *Numb) setNomerTogel(nomerTogel int) {
	u.NoTogel = nomerTogel
}

// cara membuat struct ke3
type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	p := Person{name: name, age: age}
	return &p
}

func (u *Person) changeName(newName string) {
	u.name = newName
}

func main() {
	var alamat1 = Address{3, "rusman", "jalan mahakam 1 No 3"}
	alamat1.setName("ahmad", "jalan ahamad yani no 7")

	fmt.Println(alamat1.Name())
	fmt.Println(alamat1.alamat)

	//kita bisa langsung membuat structnya dan langsung memberi nilai juga wkwkwk
	data := struct {
		Number int
		Text   string
	}{42, "Hello world"}
	fmt.Printf("%+v\n", data.Number)

	//implement struct 1
	var notogel = Numb{1, 2938}
	notogel.setNomerTogel(6547)
	fmt.Printf("%+v\n", notogel.NoTogel)

	//implement struct 2
	fmt.Println(newPerson("habibie"))
	var p1 = Person{name: "bagas", age: 90}
	p1.changeName("ahmad")
	p2 := newPerson("rahmat", 60)
	fmt.Println(p2.age)
}
