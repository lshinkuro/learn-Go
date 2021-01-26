package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 10
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("number type is  :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("number value is :", reflectValue.Int())
	}
	//pengembaliann nilai bisa menggunakan interfacd kosong
	fmt.Println(reflectValue.Interface().(int))
}
