package lesson_1

import (
	"fmt"
	"math"
)

// Напиши программу, которая запрашивает у пользователя имя и возраст, а затем выводит их.
func Task_1() {
	var name string
	var age int
	fmt.Println("Enter name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter age: ")
	fmt.Scanln(&age)
	fmt.Println("Name: ", name, ", age: ", age)
}

// Объяви три переменные разных типов (int, float64, string), присвой им значения и выведи их на экран.
func Task_2() {
	var value int = 7
	var new_value float64 = 1.7777
	var new_str string = "Aloha"

	fmt.Println(value, new_value, new_str)
}

// Создай константу pi = 3.14 и выведи ее значение.
func Task_3() {
	const pi = 3.14

	fmt.Println(pi)
}

// Напиши программу, которая запрашивает у пользователя два числа и выводит их сумму.
func Task_4() {
	var first_number int
	var second_number int
	fmt.Println("Введите первое число: ")
	fmt.Scanln(&first_number)
	fmt.Println("Введите второе число: ")
	fmt.Scanln(&second_number)
	fmt.Println("Cумма: ", first_number+second_number)

}

// Запроси у пользователя число и определи, является ли оно четным.
func Task_5() {
	var number int
	fmt.Println("Введите  число: ")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}

// Напиши программу, которая считывает строку и выводит ее длину.
func Task_6() {
	var new_str string
	fmt.Println("Введите строку:")
	fmt.Scanln(&new_str)
	fmt.Println("Длина строки: ", len(new_str))
}

// Запроси у пользователя радиус круга и вычисли его площадь (S = π * r^2).
func Task_7() {
	const pi = 3.14
	var radius float64
	fmt.Println("Введите радиус круга: ")
	fmt.Scanln(&radius)
	fmt.Println("Площадь окружности составляет: ", math.Pow(radius, 2)*pi)

}
