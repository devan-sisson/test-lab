package main

import "database/sql"


type PersonService struct {
	Db sql.DB
}

type Person struct {}


func (PersonService) GetOne(id string) (Person, error) {
	//do the stuff

	return Person{}, nil
}