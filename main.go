package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if id := params["id"]; id == item.ID {
			// <NewDecoder> & <Encode> se utilizan para enviar datos
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// if the <item> that not exists return an <empty object>
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var person Person
	// para recibir la informacion se utilizan <NewDecoder> & <Decode>
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if id := params["id"]; id == item.ID {
			people = append(people[:index], people[index+1:]...)
			fmt.Println(people)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	people = createPeople()
	router := mux.NewRouter()

	// endpoints
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func createPeople() []Person {
	var people []Person
	for i := 0; i <= 10; i++ {
		id := strconv.Itoa(i)
		people = append(people, Person{
			ID:        id,
			FirstName: "PeopleName" + id,
			LastName:  "PeopleLastName" + id,
			Address: &Address{
				City:  "City" + id,
				State: "State" + id,
			},
		})
	}
	return people
}
