package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Fullname string `json:"Name"`
	Age      int
	Address  string
}

func main() {
	// var jsonString = `{"Name":"agung","Age":23,"Address":"bandung"}`
	// var jsonParse = []byte(jsonString)

	// var user User

	// err := json.Unmarshal(jsonParse, &user)

	// fmt.Println(user.Fullname)
	// fmt.Println(user.Age)
	// fmt.Println(user.Address)

	//data bisa juga di tampung oleh map[string]interface{}

	// var user map[string]interface{}

	// err := json.Unmarshal(jsonParse, &user)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(user["Name"])
	// fmt.Println(user["Age"])

	//***data decode bisa di tampung di interface tapi hasil outputnya harus di casting supaya bisa di tampilkan

	// var user interface{}

	// err := json.Unmarshal(jsonParse, &user)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// var decodeData = user.(map[string]interface{})

	// fmt.Println(decodeData["Name"])
	// fmt.Println(decodeData["Age"])

	//***data decode ditampung oleh array

	var jsonString = `[
		{"Name":"agung","Age":23,"Address":"bandung"},
		{"Name":"ari","Age":22,"Address":"subang"}
		]`

	var user []User

	err := json.Unmarshal([]byte(jsonString), &user)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < len(user); i++ {
		fmt.Println(user[i].Fullname)
		fmt.Println(user[i].Age)
		fmt.Println(user[i].Address)

	}

}
