package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan int, 2) //jumlah buffer 3 karena index mulai 0

	go func() {
		for {
			i := <-message
			fmt.Println("receive data :", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data :", i)
		message <- i
	}

	var who string
	fmt.Scanln(&who)
}
