package iterator

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade string
}

func NewStudent(name string, age int, grade string) *Student {
	return &Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
}

func (s *Student) String() string {
	return fmt.Sprintf("Name: %s Age: %d Grade:%s", s.Name, s.Age, s.Grade)
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) GetAge() int {
	return s.Age
}
func (s *Student) GetGrade() string {
	return s.Grade
}
