package main

import (
	"fmt"
	"runtime"
)

func AverageNum(numbers []int, ch chan float64) {
	var sum = 0
	for _, i := range numbers {
		sum += i
	}
	ch <- float64(sum) / float64(len(numbers))
}

func MaxNumber(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func main() {

	runtime.GOMAXPROCS(2)

	var numbers = []int{1, 3, 5, 4, 6, 7, 2, 9, 3}

	var ch1 = make(chan float64)
	go AverageNum(numbers, ch1)

	var ch2 = make(chan int)
	go MaxNumber(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Println("max number is ", max)
		}
	}

}
