package main

import (
	"fmt"
)

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "wick",
	height:      182.5,
	age:         26,
	isGraduated: false,
	hobbies:     []string{"eating", "sleeping"},
}

func main() {
	fmt.Printf("%b \n", data.age)
	// %c memformat data numerik yang merupakan kode unicode, menjadi bentuk string karakter unicode-nya.
	fmt.Printf("%c \n", 1400)
	// %d  memformat data numerik, menjadi bentuk string numerik berbasis 10 (basis bilangan yang kita gunakan).
	fmt.Printf("%d \n", data.age)
	// $e || %E memformat data numerik desimal ke dalam bentuk notasi numerik standar Scientific notation.
	fmt.Printf(" %e \n", data.height)
	fmt.Printf("%E \n", data.height)
	// $f || %F  memformat data numerik desimal, dengan lebar desimal bisa ditentukan. Secara default lebar digit desimal adalah 6 digit.
	fmt.Printf("%f \n", data.height)
	fmt.Printf("%.9f \n", data.height)
	fmt.Printf("%.2f \n", data.height)
	fmt.Printf("%.f \n", data.height)
	// %g || %G memformat data numerik desimal, dengan lebar desimal bisa ditentukan. Lebar kapasitasnya sangat besar,
	fmt.Printf("%.g \n", 0.23480000000000002983)
	fmt.Printf("%G \n", 0.23480000000000002983)
	// %o  memformat data numerik, menjadi bentuk string numerik berbasis 8 (oktal).
	fmt.Printf("%o \n", data.age)
	// %p mengembalikan alamat pointer referensi variabel-nya.Alamat pointer dituliskan dalam bentuk numerik berbasis 16 dengan prefix 0x
	fmt.Printf("%p \n", &data.age)
	// %q  escape string. Meskipun string yang dipakai menggunakan literal \ akan tetap di-escape.
	fmt.Printf("%q \n", `"\\ name \ hai"`)
	// %s  memformat data string
	fmt.Printf("%s \n", data.name)
	// %t untuk menanpilkan isi data boolean
	fmt.Printf("%t \n", data.isGraduated)
	// %T untuk menampilkan type variabel format
	fmt.Printf("%T \n", data.name)
	// %v %+v %#+v di gunakan memformat data
	fmt.Printf("%v \n", data)
	fmt.Printf("%+v \n", data)
	fmt.Printf("%#+v \n", data)
	// %x || %X digunakan untuk memformat menjadi 16 hexsa desimal bentuk string
	fmt.Printf("%x \n", data.height)
	//%% digunakan untuk mennulis karakter string % pada console
	fmt.Printf("% %% \n")
}
