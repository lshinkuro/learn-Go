package main

import (
	"fmt"
)

func main() {
	var pow float64 = 2 << 3 << 3
	var x *float64 = &pow
	fmt.Println(*x)
}
