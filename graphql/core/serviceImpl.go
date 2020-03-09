package core

import (
	"github.com/go-kit/kit/log"
)

type PersonServImpl struct {
	personRepo PersonRepository
	logger     log.Logger
}

func CreatePersonService(personRepo PersonRepository, log log.Logger) *PersonServImpl {
	return &PersonServImpl{
		personRepo: personRepo,
		logger:     log,
	}
}

func (ps *PersonServImpl) AddPerson(person Person) (string, error) {
	logger := log.With(ps.logger, "method", "AddPerson")
	result, err := ps.personRepo.SavePerson(person)
	if err != nil {
		return "failed", err
	}
	logger.Log("Person Creeated")
	return result, nil

}

func (ps *PersonServImpl) FetchPersonById(pid string) Person {
	return ps.personRepo.GetPersonById(pid)
}

func (ps *PersonServImpl) FetchAllPersons() []Person {
	return ps.personRepo.GetPersons()
}
