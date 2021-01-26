package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Println("masukan input :")
	fmt.Scanln(&input)

	var number int
	var err error

	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println("your input is number ,", number)
	} else {
		fmt.Println("your input is not number")
		fmt.Println(err.Error())
	}

}
