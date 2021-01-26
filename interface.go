package main

import (
	"fmt"
)

//ini adalah cara untuk mebuat struct
type Address struct {
	City, Province, Country string
	Id                      int
}

func ChangeCity(address *Address, id int) {
	address.City = "malang"
	address.Id = id
}

func main() {
	// ini adalah cara untuk memberi nilai pada struct
	var address = Address{"bandung", "jawa barat", "indonesia", 2}
	//ini adalah cara untuk membuat pointer

	var address1 *Address = &address
	// pointer adalah cara mengubah nilai variabel yang merujuk pada satu nilai yang di inisialisasi
	address1.City = "indramayu"

	fmt.Println(address.City)
	fmt.Println(address1.City)

	var addres2 *Address = &address
	ChangeCity(addres2, 9)
	fmt.Println(addres2.Id)

}

// type Address struct {
// 	City, Province, Country string
// }

// func ChangeCountry(address *Address) {
// 	address.Country = "malaysia"
// }

// func main() {
// 	var address1 Address = Address{"jakarta", "jaksel", "indonesia"}
// 	var address2 *Address = &address1
// 	var address3 *Address = &address1

// 	address2.City = "bandung"

// 	*address2 = Address{"malang", "bumiayu", ""}

// 	fmt.Println(address1)
// 	fmt.Println(address2.City)
// 	fmt.Println(address3)

// 	var alamat Address = Address{
// 		City:     "malang",
// 		Province: "sumberrejo",
// 		Country:  "",
// 	}

// 	var alamatBaru *Address = &alamat
// 	ChangeCountry(alamatBaru)
// 	fmt.Println(alamatBaru)

// }
