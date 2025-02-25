package lesson_8

import "fmt"

type Speaker interface {
	Speak() string
}

type Human struct {
	Name string
}
type Dog struct {
	Breed string
}

func (h Human) Speak() string {
	return "Hi, I'm " + h.Name
}

func (d Dog) Speak() string {
	return "Hi, I'm " + d.Breed
}
func Introduce(s Speaker) {
	fmt.Println(s.Speak())
}

func Lesson_8() {
	d := Dog{"Labrador"}
	h := Human{"Ivan"}

	Introduce(d)
	Introduce(h)
}
