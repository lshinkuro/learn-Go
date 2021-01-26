package main

import (
	"fmt"
	"strings"
)

func filter(message []string, callback func(string) bool) []string {
	var result []string
	for _, each := range message {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
func jon(s string) bool {
	return strings.Contains(s, "o")
}

func main() {
	var data = []string{"banu", "robi", "ikhsan", "akmal", "foden"}

	var dataContains = filter(data, jon)

	var dataLen = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println(data)
	fmt.Println(dataContains)
	fmt.Println(dataLen)

}
