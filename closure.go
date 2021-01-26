package main

import (
	"fmt"
	"sort"
)

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case i > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var sorting = func(nlist []int) []int {

		for _, y := range nlist {
			fmt.Println(y)
		}
		sort.Ints(nlist)
		return nlist
	}

	var numbers = []int{9, 4, 5, 3, 9, 2, 4, 1}

	var min, max = getMinMax(numbers)

	fmt.Println(numbers, min, max)

	var sort = sorting(numbers)

	fmt.Println(sort)

}
