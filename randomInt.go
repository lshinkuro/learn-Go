package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//fungsi ini du gunakan untuk mencari random string baru
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(2))
}
