package structs

import "errors"

type Person struct {
	firstName  string //private fields
	lastName   string
	age        int
	middleName []string
}

// constructor
func NewPerson(fn, ln string, a int, mn []string) Person {
	//validation logic and if it fails then panic
	return Person{
		firstName:  fn,
		lastName:   ln,
		age:        a,
		middleName: mn,
	}
}

func (p Person) Equals(p1 Person) bool {
	if len(p.middleName) == len(p1.middleName) {
		return true
	}
	return false
}

// receiver functions
func (p Person) Walk() string {
	return p.firstName + p.lastName + " likes walking"
}

// pointer based receiver function
func (p *Person) SetFirstName(f string) error {
	if len(f) < 3 {
		return errors.New("invalid firstname")
	}
	p.firstName = f
	return nil
}

// Getter function
func (p *Person) FirstName() string {
	return p.firstName
}
