package core

import (
	"fmt"
)

type PersonServImpl struct {
	personRepo PersonRepository
}

func CreatePersonService(personRepo PersonRepository) *PersonServImpl {
	return &PersonServImpl{personRepo: personRepo}
}

func (ps *PersonServImpl) SavePerson(person Person) (string, error) {
	return ps.personRepo.SavePerson(person)
}

func (ps *PersonServImpl) GetPersonById(pid string) Person {
	fmt.Println("in service impl layer")
	return ps.personRepo.GetPersonById(pid)
}

func (ps *PersonServImpl) GetPersons() []Person {
	return ps.personRepo.GetPersons()
}
