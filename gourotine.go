package main

import (
	"fmt"
	"runtime"
	"time"
)

func print(till int, message string) {
	for i := 0; i <= till; i++ {
		fmt.Println(i+1, "message is :", message)
	}
}

func main() {
	//run time gunanya untuk menentukan jumlah core cpu yang akan kita gunakan
	runtime.GOMAXPROCS(2)

	fmt.Println("halo semuanya belajar gourotine")

	go print(9, "hallo")
	print(4, "semuanya")
	go func(message string) {
		for i := 0; i <= 4; i++ {
			fmt.Println(message)
		}
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	var input1, input2 string
	fmt.Scanln(&input1, &input2)
	fmt.Println(input1)
	fmt.Println(input2)

}
