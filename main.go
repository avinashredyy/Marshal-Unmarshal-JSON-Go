package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//Set content Type and status
	w.Header().Set("Contetnt-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Open our jsonFile which mimics the data recieved from Mongo
	jsonFile, err := os.Open("data.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array. We do this so that we can unmarshal
	//the contents of byteValue based on the definition of Users struct
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our jsonFile's content into 'users' which
	//we defined above. Only those properties will be accepted which are defined in Users struct
	json.Unmarshal(byteValue, &users)
	// fmt.Println(users.Users)
	// usersType := fmt.Sprintf("%T", users.Users)
	// fmt.Println(usersType)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	// for i := 0; i < len(users.Users); i++ {
	// 	fmt.Println("User Type: " + users.Users[i].Type)
	// 	fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
	// 	fmt.Println("User Name: " + users.Users[i].Name)
	// 	fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	// }
	//we marshal back the json content based on the Users struct defination.
	resp, err := json.Marshal(users.Users)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}
