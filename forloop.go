package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

}
