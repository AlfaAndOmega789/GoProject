package chapter_2

import "fmt"

// 1. Нулевое значение (Легкий)
// Напишите программу, которая объявляет переменные разных встроенных типов (int, float64, bool, string) без явного присваивания значений и затем выводит их на экран.
func ZeroValue() {
	var num1 int
	var num2 float64
	var value bool
	var str string

	fmt.Println(num1, num2, value, str)
}

//
