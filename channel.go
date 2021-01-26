package main

import (
	"fmt"
	"runtime"
)

//chanel adalah pertukaran data di antara goroutine bisa mengirim data bisa menerima data

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var SaytoHello = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go SaytoHello("jhon")
	go SaytoHello("kholis")
	go SaytoHello("abdul")

	var messages1 = <-messages
	fmt.Println(messages1)
	var messages2 = <-messages
	fmt.Println(messages2)
	var messages3 = <-messages
	fmt.Println(messages3)

}
