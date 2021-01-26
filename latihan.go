package main

import (
	"fmt"
	"strings"
)

type person struct {
	name         string
	grade, nilai int
}

func filterName(ctx []person, char string) []string {
	var data []string
	for _, e := range ctx {
		if filtered := strings.Contains(e.name, char); filtered {
			data = append(data, e.name)
		}
	}
	return data
}

func maxNumber(ctx []person) int {
	var max = ctx[0].nilai
	for _, e := range ctx {
		if max < e.nilai {
			max = e.nilai
		}
	}
	return max
}

func main() {
	var data = []person{
		{name: "agung", grade: 3, nilai: 9},
		{name: "obug", grade: 3, nilai: 7},
	}

	P1 := filterName(data, "g")
	fmt.Println(P1)
	Max := maxNumber(data)
	fmt.Println(Max)
}
