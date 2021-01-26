package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Grade int
}

func (s *Student) GetPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)
	
	
	//untuk mengecek apakah tipe datanya pointer
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
    
	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("Name is :", reflectType.Field(i).Name)
		fmt.Println("type is :", reflectType.Field(i).Type)
		fmt.Println("value is :", reflectValue.Field(i).Interface())
		fmt.Println(" ")
	}
}

func main() {
	var S1 = Student{"riki", 30}

	S1.GetPropertyInfo()
}
