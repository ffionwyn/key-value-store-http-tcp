package store

import (
	"testing"
)

//Test for adding someone to storage.
func TestAddToStorage(t *testing.T) {
	name := "ffion"
	location := "wrexham"
	age := "28"
	err := AddToStorage(name, location, age)
	if err != nil {
		t.Error(`err := AddToStorage(name, location, age) not equal to nil`)
	}
	exists := CheckPerson(name)
	if !exists {
		t.Error(`error: person hasnt been added`)
	}
}

//Test for incorrectly trying to add someone to storage with empty string age.
func TestAddToStorageEmptyInput(t *testing.T) {
	name := "ffion"
	location := "wrexham"
	age := ""
	err := AddToStorage(name, location, age)
	if err == nil {
		t.Error(`no error for invalid input`)
	}
}

//Test for incorrectly trying to add someone to storage with incorrect data type.
func TestAddToStorageInvalidInput(t *testing.T) {
	name := "ffion"
	location := "wrexham"
	age := "not a number"
	err := AddToStorage(name, location, age)
	if err == nil {
		t.Error(`no error for invalid input`)
	}
}

//Test for trying to delete someone who isnt in the storage.
func TestDeletePersonThatDoesNotExist(t *testing.T) {
	name := "izzy"
	err := DeletePerson(name)
	if err == nil {
		t.Error(`no error for invalid input`)
	}
}

//Test for deleting an existing person from storage.
func TestDeletePersonThatDoesExist(t *testing.T) {
	AddToStorage("ffion", "wrexham", "28")
	name := "ffion"
	err := DeletePerson(name)
	if err != nil {
		t.Error(`unexpected error`)
	}
}

//Test for printing out persons location via their name.
func TestGetPersonLocation(t *testing.T) {
	AddToStorage("ffion", "wrexham", "28")
	name := "ffion"
	location, age, err := GetPersonLocation(name)
	if err != nil {
		t.Error(`error failed to get person location`)
	}
	if location != "wrexham" {
		t.Error(`error location invalid input`)
	}
	if age != 28 {
		t.Error(`error age invalid input`)
	}
}

//Test for updating someones location.
func TestUpdatePersonStorage(t *testing.T) {
	AddToStorage("ffion", "wrexham", "28")
	UpdatePersonStorage("ffion", "manchester", "28")
	name := "ffion"
	location, age, err := GetPersonLocation(name)
	if err != nil {
		t.Error(`error failed to update person`)
	}
	if location != "manchester" {
		t.Error(`error invalid location input`)
	}
	if age != 28 {
		t.Error(`error invalid age input`)
	}
}

//Test for trying to update someones location that doesnt exist.
func TestUpdatePersonStorageMissingPerson(t *testing.T) {
	err := UpdatePersonStorage("minnie", "cardiff", "21")
	if err == nil {
		t.Error(`no error for invalid input`)
	}
}

//Test for empty string on validation.
func TestValidationInput(t *testing.T) {
	name := "ffion"
	location := "wrexham"
	age := ""
	err := validateInput(name, location, age)
	if err == nil {
		t.Error(`no error for invalid input`)
	}
}