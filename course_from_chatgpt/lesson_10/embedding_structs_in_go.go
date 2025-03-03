package lesson_10

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	Job     string
	Salary  int
	Company string
}

func (p Person) Speak() {
	fmt.Println("Hi, my name is", p.Name)
}
func Lesson_10() {
	e1 := Employee{
		Person: Person{"Ivan", 26},
		Job:    "Programmer",
		Salary: 4000,
	}
	fmt.Println(e1.Name, e1.Age)
	fmt.Println(e1.Job, e1.Salary)

	e2 := Employee{
		Person:  Person{Name: "Ivan"},
		Company: "Google",
	}
	e2.Speak()
}
func (e Employee) Speak() {
	fmt.Println("I working in", e.Company)
}
