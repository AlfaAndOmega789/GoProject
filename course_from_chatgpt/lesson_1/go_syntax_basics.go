package lesson_1

import "fmt"

func Lesson_1() {
	fmt.Println("Hello, world!")

	//Variables and constants
	var name string = "Alice" // явное указание типа
	var age int = 25
	weight := 68.5 // короткое определение типа(тип определяется автоматически)

	fmt.Println("Имя: ", name)
	fmt.Println("Возраст: ", age)
	fmt.Println("Вес: ", weight)

	const pi = 3.14
	fmt.Println("Pi: ", pi)

	//Data types
	var number int = 1
	var other_number float64 = 1.54
	var s string = "Golang"
	var value bool = true
	var char rune = 'A'

	fmt.Println(number, other_number, s, value, char)

	//Data entry
	var new_name string
	fmt.Print("Введите имя: ")
	fmt.Scanln(&new_name) //считывает строку

	fmt.Println("Привет: ", new_name)
}
