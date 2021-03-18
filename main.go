package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"Name,omitempty"`
	Address     string `json:"Address,omitempty"`
	Phonenumber string `json:"Phonenumber,omitempty"`
}
type Mobile struct {
	ID           string `json:"ID,omitempty"`
	Model        string `json:"Model,omitempty"`
	manufacturer string `json:"manufacturer,omitempty"`
	Costprice    string `json:"Cost-price,omitempty"`
	Sellingprice string `json:"Selling-price,omitempty"`
}

var people []Customer
var Owner []Mobile

func GetPersonByid(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func GetPersonList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, p := range people {
		if p.ID == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("Customer not found")
}
func AddPerson(w http.ResponseWriter, r *http.Request) {
	var Customer Customer
	_ = json.NewDecoder(r.Body).Decode(&Customer)
	people = append(people, Customer)
	json.NewEncoder(w).Encode(Customer)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, p := range people {
		if p.ID == params["id"] {
			copy(people[i:], people[i+1:])
			people = people[:len(people)-1]
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}
func GetMobileByid(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Owner)
}
func GetMobileList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, p := range Owner {
		if p.ID == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("Mobile not found")
}
func AddMobile(w http.ResponseWriter, r *http.Request) {
	var Mobile Mobile
	_ = json.NewDecoder(r.Body).Decode(&Mobile)
	Owner = append(Owner, Mobile)
	json.NewEncoder(w).Encode(Mobile)
}
func DeleteMobile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, p := range Owner {
		if p.ID == params["id"] {
			copy(Owner[i:], Owner[i+1:])
			Owner = Owner[:len(Owner)-1]
			break
		}
	}
	json.NewEncoder(w).Encode(Owner)
}

func main() {
	fmt.Println("Starting server on port 8000...")

	router := mux.NewRouter()
	people = append(people, Customer{ID: "1", Name: "Bruce", Address: "Jp road", Phonenumber: "012345678"})
	people = append(people, Customer{ID: "2", Name: "Clark", Address: "Gandhi road", Phonenumber: "9876543210"})
	router.HandleFunc("/people", GetPersonByid).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonList).Methods("GET")
	router.HandleFunc("/people", AddPerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	Owner = append(Owner, Mobile{ID: "1", Model: "Nokia", manufacturer: "2013", Costprice: "3000", Sellingprice: "4000"})
	Owner = append(Owner, Mobile{ID: "2", Model: "Mi", manufacturer: "2014", Costprice: "5000", Sellingprice: "7000"})
	router.HandleFunc("/Mobile", GetMobileByid).Methods("GET")
	router.HandleFunc("/Mobile/{id}", GetMobileList).Methods("GET")
	router.HandleFunc("/Mobile", AddMobile).Methods("POST")
	router.HandleFunc("/Mobile/{id}", DeleteMobile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
