package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%+v \n", now)
	//cara passing string ke format time
	var value string
	value = "02 Sep 15 08:00 WIB"
	var date, _ = time.Parse(time.RFC822, value)
	fmt.Println(date.String())
	// 2015-09-02 08:00:00 +0700 WIB

	//cara format dari time ke string
	var dateS1 = date.Format(time.RFC822)
	fmt.Println("data string is ", dateS1)
}
