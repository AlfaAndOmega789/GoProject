package lesson_7

import "fmt"

func Lesson_7() {
	p := Person{"Ivan", 26}
	fmt.Println(p.Age, p.Name)

	upgradeAge(&p, 18)
	fmt.Println(p.Age, p.Name)

	p.SayHello()
}

type Person struct {
	Name string
	Age  int
}

func upgradeAge(p *Person, newAge int) {
	p.Age = newAge

}
func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}
