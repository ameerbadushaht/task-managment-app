package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users Users

type PageData struct {
	Uname string
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	tmpl := template.Must(template.ParseFiles("logged.html"))
	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].Username == username && users.Users[i].Password == password {
			data := PageData{Uname: string(users.Users[i].Username)}
			tmpl.Execute(w, data)
		}

	}

}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("user.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example

	http.HandleFunc("/api", handleRequest)
	// add our articles route and map it to our
	http.ListenAndServe(":10000", nil)
}
