package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var names = []string{"ardi", "hanif"}
// 	printMessage("hallo", names)

// }

// func printMessage(message string, arr []string) {
// 	var nameString = strings.Join(arr, " ")
// 	fmt.Println(message, nameString)
// }

import (
	"fmt"
	"math"
)

// func main() {
// 	rand.Seed(time.Now().Unix())

// 	var randomValue int
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number :", randomValue)
// }

// func randomWithRange(min, max int) int {
// 	var value = rand.Int()%(max-min+1) + min
// 	return value
// }

type lingkaran struct {
	diameter float64
}

func (r *lingkaran) hitung() (luas float64, keliling float64) {
	luas = math.Pi * math.Pow(r.diameter/2, 2)
	keliling = math.Pi * r.diameter

	return

}

func main() {
	var L1 = lingkaran{diameter: 7}
	var area, keliling = L1.hitung()
	fmt.Println("laus dan kelilingnya adalah :", area, keliling)

}
