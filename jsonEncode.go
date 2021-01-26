package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Fullname string `json:"Name"`
	Age      int
	Address  string `json  :"Alamat"`
	Gender   bool
}

func main() {
	// var jsonString = `{"Name":"Ajung","Age":12,"Alamat":"Cirebon","Gender":false}`
	// var jsonParse = []byte(jsonString)
	// var data User

	// err := json.Unmarshal(jsonParse, &data)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(data.Gender)

	//*** encode adalah mengubah data string bisa berupa slice map atau pun variabel biasa untuk di rubah ke json

	var jsonString = []User{
		{Fullname: "Ajung", Age: 22, Address: "Cirebon", Gender: false},
		{Fullname: "Ahdan", Age: 22, Address: "Sumedang", Gender: false},
	}

	var EncodeData, err = json.Marshal(jsonString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data := string(EncodeData)

	fmt.Println("sudah string ini", data)

	//di decode lagi supaya jsonya bisa jadi object

	var user []User

	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(user[0].Fullname)
}
