package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("can't be empty")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("error occured", r)
	} else {
		fmt.Println("application run perfectly")
	}
}

//recover digunakan untuk menangani agar error panic tidak muncul di console

func main() {
	defer catch()

	var input string
	fmt.Print("masukan nama mu : ")
	fmt.Scanln(&input)

	if valid, err := validate(input); valid {
		fmt.Println("halo :", input)
	} else {
		//penggunaan panic memberhentikan semua proses jadi ketika ada error maka panic akan
		//bekerja dan baris kode setelahnya tidak akan dijalankan kecuali di defer terlebih dahulu
		panic(err.Error())
		fmt.Println("end")
	}
}
