package core

type PersonService interface {
	SavePerson(person Person) (string, error)
	GetPersonById(pid string) Person
	GetPersons() []Person
}
