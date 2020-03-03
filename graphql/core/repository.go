package core

type PersonRepository interface {
	SavePerson(person Person) (string, error)
	GetPersonById(pid string) Person
	GetPersons() []Person
}
