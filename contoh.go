package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	go func() {
		data := fmt.Sprintf("halo semuanya")
		messages <- data
	}()

	fmt.Println(<-messages)
}
