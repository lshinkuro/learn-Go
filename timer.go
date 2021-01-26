package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(time.Second * 4)
	fmt.Println("hallo semuanya")
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("halo ", i)
	}
}
