package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Printf("years is %v and month is %v \n", now.Year(), now.Month())

	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())
	// 2015-09-02 08:04:00 +0000 UTC

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())
	// 2015-09-02 00:00:00 +0700 WIB
}

// Data string bisa dikonversi menjadi time.Time dengan memanfaatkan time.Parse. Fungsi ini membutuhkan 2 parameter:

// Parameter ke-1 adalah layout format dari data waktu yang akan diparsing.
// Parameter ke-2 adalah data string yang ingin diparsing.