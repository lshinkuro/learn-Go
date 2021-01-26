package main

import (
	"fmt"
	"runtime"
)

func PrintMessage(who chan string) {
	fmt.Println(<-who)
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _, each := range []string{"heru", "malik", "bagas"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		PrintMessage(messages)
	}

}
