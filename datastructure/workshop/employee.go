package testproj

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func InitEmployee(firstName, lastlastName string, id int) *Employee {
	return &Employee{
		firstName: firstName,
		lastName:  lastlastName,
		id:        id,
	}
}

func (e *Employee) getFirstName() string {
	return e.firstName
}

func (e *Employee) getLastName() string {
	return e.lastName
}

func (e *Employee) getID() int {
	return e.id
}

func (e *Employee) setFirstName(name string) {
	e.firstName = name
}

func (e *Employee) setLastName(lname string) {
	e.firstName = lname
}

func (e *Employee) setID(id int) {
	e.id = id
}

func (e *Employee) toString() string {
	return fmt.Sprintf("Employee{ firstName=%s, lastName=%s, Id=%d}", e.firstName, e.lastName, e.id)
}
