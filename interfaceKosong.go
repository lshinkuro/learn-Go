package main

import (
	f "fmt"
	"strings"
)

func main() {
	f.Println("belajar interface kosong")
	//cara membuat ionterface kosong
	var secret interface{}

	secret = 8

	f.Println(secret)

	secret = []string{"anjing", "bangsat", "babi"}

	f.Println(secret)

	//interface sebagai tipe data dalam membuat map

	var data = map[string]interface{}{
		"name":      "agung",
		"grade":     8,
		"breakfast": []string{"apple", "manggo", "banana"},
	}
	f.Println(data)

	//untuk melakukan perhitungan dalam interface di butuhkan casting variable agar bisa memunculkan nilai aslinya

	var secret1 interface{}

	secret1 = 2
	//ini adalah methode casting

	var number = secret1.(int) * 18

	f.Println(number)

	secret1 = []string{"apple", "mango", "banana"}

	var gruits = strings.Join(secret1.([]string), ",")

	f.Println(gruits, "is my favorite fruits")

	//casting bisa untuk struck juga

	type person struct {
		name string
		age  int
	}

	var secret2 interface{} = &person{"agung", 19}

	var name = secret2.(*person).name
	//ingat pointer adalah tipe data jadi bisa di casting
	f.Println(name)

	//di contoh ini saya akan menggabungkan slice ,map,dan interface kosong

	var secret3 = []map[string]interface{}{
		{"name": "ahmad", "age": 34},
		{"name": "robi", "age": 23},
	}

	for _, each := range secret3 {
		f.Println(each["name"])
	}

}

// Interface kosong atau empty interface yang dinotasikan dengan interface{}
// merupakan tipe data yang sangat spesial. Variabel bertipe ini bisa menampung segala jenis data,
// bahkan array, pointer, apapun. Tipe data dengan konsep ini biasa disebut dengan dynamic typing.

// interface seperti yang kita tau, digunakan untuk pembuatan interface. Tetapi ketika ditambahkan kurung kurawal ({})
//  di belakang-nya (menjadi interface{}), maka kegunaannya akan berubah, yaitu sebagai tipe data.
