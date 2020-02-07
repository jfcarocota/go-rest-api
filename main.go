package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string  `json:"id,omitempty"`
	FirstName string  `json:"forstname,omitempty"`
	LastName  string  `json:"lastname,omitempty"`
	Adress    *Adress `json:"adress,omitempty"`
}

type Adress struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndPoint(W http.ResponseWriter, req *http.Request) {
	json.NewEncoder(W).Encode(people)
}

func GetPersoneEndPoint(W http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(W).Encode(item)
			return
		}
	}
	json.NewEncoder(W).Encode(&Person{})
}

func CreatePersonEndPoint(W http.ResponseWriter, req *http.Request) {

}

func DeletePersonEndPoint(W http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{
		ID: "1", FirstName: "Jesus",
		LastName: "Caro",
		Adress: &Adress{
			City: "Obregon", State: "Sonora"}})
	people = append(people, Person{
		ID:        "2",
		FirstName: "Jhon",
		LastName:  "Doe",
		Adress: &Adress{
			City: "Las Vegas", State: "Nevada"}})

	//endoint
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersoneEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
