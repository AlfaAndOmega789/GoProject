package lesson_6

import "fmt"

func Lesson_6() {
	a := 10
	p := &a // р хранит адрес переменной a

	fmt.Println("Значени, aе a:")
	fmt.Println("Адрес а:", p)
	fmt.Println("Значение через указатель:", *p)

	*p = 20
	fmt.Println("Новое значение a:", a)

}
