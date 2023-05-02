package server

import (
	"fmt"
	"log"
	"net/http"
	"tcp/store"
)

//Routes the request depending on the method. 
func PeopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		addPerson(w, r)
	case "GET":
		getPerson(w, r)
	case "DELETE":
		DeletePerson(w, r)
	case "PATCH":
		updatePerson(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Sorry, only GET/DELETE/PATCH/POST methods are supported.")
	}
}

//Function for printing every person in the store.
func PrintStorage(w http.ResponseWriter, r *http.Request) {
	allPeople := store.PrintPersonStorage()
	fmt.Fprint(w, allPeople)
	w.WriteHeader(http.StatusOK)
}

//Function for adding someone new in the store, needs name location and age.
func addPerson(w http.ResponseWriter, r *http.Request) {
	inputtedName := r.URL.Query().Get("name")
	inputtedLocation := r.URL.Query().Get("location")
	inputtedAge := r.URL.Query().Get("age")
	err := store.AddToStorage(inputtedName, inputtedLocation, inputtedAge)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

//Function for getting someones location via name.
func getPerson(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	location, age, err := store.GetPersonLocation(name)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	message := fmt.Sprintf("location: %s, age %d", location, age)
	w.Write([]byte(message))
	w.WriteHeader(http.StatusOK)
}

//Function for deleting somebody who is in the store via name.
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	err := store.DeletePerson(name)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//Function for updating name location or age of person in the store.
func updatePerson(w http.ResponseWriter, r *http.Request) {
	inputtedName := r.URL.Query().Get("name")
	inputtedLocation := r.URL.Query().Get("location")
	inputtedAge := r.URL.Query().Get("age")
	err := store.UpdatePersonStorage(inputtedName, inputtedLocation, inputtedAge)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}