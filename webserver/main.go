package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "halo apa kabar !!")
}

type User struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", index)

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = []User{
			{Name: "agung", Age: 12},
			{Name: "bagas", Age: 23},
		}
		t, err := template.ParseFiles("template.html")

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(data[0].Name)

		t.Execute(w, data[0])

	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		var jsonString = []User{
			{Name: "Ajung", Age: 22},
			{Name: "Ahdan", Age: 22},
		}
		var EncodeData, err = json.Marshal(jsonString)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data := string(EncodeData)

		var user []User

		err = json.Unmarshal([]byte(data), &user)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Fprintln(w, user[0].Name)
	})

	//simple web rest API

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var data = []User{
			{Name: "Ajung", Age: 22},
			{Name: "Ahdan", Age: 29},
		}

		if r.Method == "POST" {
			var EncodeData, err = json.Marshal(data)
			if err != nil {
				fmt.Println(err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(EncodeData)
			return

		}

		http.Error(w, "", http.StatusBadRequest)
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data = []User{
			{Name: "Ajung", Age: 22},
			{Name: "Ahdan", Age: 22},
		}

		if r.Method == "POST" {
			var id = r.FormValue("id")
			var result []byte
			var err error

			for _, each := range data {
				if each.Name == id {
					result, err = json.Marshal(each)

					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					w.Write(result)
					return
				}
			}
		}
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	})

	fmt.Println("web server run on port :8080")
	http.ListenAndServe(":8080", nil)
}
