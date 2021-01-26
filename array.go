package main

import (
	"fmt"
)

func main() {
	//how to make array in  golang
	var nama = [3]string{"ahmad", "agung", "bagas"}
	fmt.Println(len(nama))

	//menampilkan array
	for i, v := range nama {
		fmt.Println(i, v)
	}

	//inisialisasi array dengan make
	var fruits = make([]string, 2)

	fruits[0] = "apple"
	fruits[1] = "manggo"

	for i := 0; i <= len(fruits)-1; i++ {
		fmt.Println(fruits[i])
	}

	//cara membuat map;
	var chicken = map[string]int{"januari": 50, "februari": 20}
	fmt.Println(chicken["januari"])
	for _, v := range chicken {
		fmt.Println(v)
	}

	// var chicken2 = *new(map[string]int)

	var chicken1 = map[string]int{"januari": 6, "februari": 49}
	//cara menghapus salah satu elemen array
	delete(chicken1, "januari")
	fmt.Println(chicken1["februari"])

	//gabungan slice dan map

	var chicken3 = []map[string]string{
		{"name": "agung", "gender": "male"},
		{"name": "rahmat", "gender": "male"},
		{"name": "bagas", "gender": "male"},
	}

	for _, v := range chicken3 {
		for _, j := range v {
			fmt.Println("halo", j)
		}
	}

}
