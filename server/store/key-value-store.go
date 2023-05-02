package store

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type Person struct {
	location string
	age      int
}

func newPerson(loc string, age int) Person {
	p := Person{
		location: loc,
		age:      age,
	}
	return p
}

var personStorage = make(map[string]Person)

//Function for deleting somebody who is in the store via name.
func DeletePerson(name string) error {
	if _, ok := personStorage[name]; !ok {
		return errors.New("person does not exist")
	}
	delete(personStorage, name)
	log.Println("Delete person successful")
	return nil
}

//Function for adding someone new in the store, needs name location and age.
func AddToStorage(name string, location string, agestring string) error {
	age, ageErr := strconv.Atoi(agestring)
	if ageErr != nil {
		return ageErr
	}
	validationErr := validateInput(name, location, agestring)
	if validationErr != nil {
		return validationErr
	}
	fmt.Println("Hello " + name + " ")
	p := newPerson(location, age)
	personStorage[name] = p
	log.Println("Added to storage successful")
	return nil
}

//Function for getting someones location via name.
func GetPersonLocation(name string) (string, int, error) {
	Person, found := personStorage[name]
	if !found {
		return "", 0, fmt.Errorf("person does not exist")
	}
	return Person.location, Person.age, nil
}

//Function for checking if a person exists in the store.
func CheckPerson(name string) bool {
	_, exists := personStorage[name]
	return exists
}

//Function for updating name location or age of person in the store.
func UpdatePersonStorage(name string, location string, agestring string) error {
	validationErr := validateInput(name, location, agestring)
	if validationErr != nil {
		return validationErr
	}
	age, err := strconv.Atoi(agestring)
	if err != nil {
		return err
	}
	_, ok := personStorage[name]
	if !ok {
		log.Print("person not in storage - failed to update")
		return fmt.Errorf("person does not exist")
	}
	p := newPerson(location, age)
	personStorage[name] = p
	log.Println("Update person successful")
	return nil
}

//Function for printing every person in the store.
func PrintPersonStorage() string {
	var message string
	for key, person := range personStorage {
		message += fmt.Sprintf("name:%s, location: %s, age %d \n", key, person.location, person.age)
	}
	return message
}

//Function for validating user input ensuring nothing is empty.
func validateInput(name string, location string, agestring string) error {
	if name == "" {
		return errors.New("missing name parameter")
	}
	if location == "" {
		return errors.New("missing location parameter")
	}
	if agestring == "" {
		return errors.New("missing age parameter")
	}
	return nil
}