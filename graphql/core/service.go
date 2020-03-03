package core

type PersonService interface {
	AddPerson(person Person) (string, error)
	FetchPersonById(pid string) Person
	FetchAllPersons() []Person
}
