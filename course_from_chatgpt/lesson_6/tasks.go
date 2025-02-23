package lesson_6

import (
	"fmt"
	"math"
)

// Создай указатель на целое число, присвой ему адрес переменной и выведи её значение через разыменование.
func Task_1() {
	number := 10

	point := &number

	fmt.Println("Значение через разыменование:", *point)

}

// Напиши функцию, которая принимает указатель на число и увеличивает его значение на 10.
func Task_2(point *int) {
	*point += 10
}

// Создай две переменные и обменяй их значениями, используя указатели.
func Task_3() {
	num1 := 11
	num2 := 22
	fmt.Println(num1, num2)

	n1 := &num1
	n2 := &num2

	num1 = *n2
	num2 = *n1
	fmt.Println(num1, num2)
}

// Напиши функцию, которая меняет местами два элемента в срезе, используя указатели.
func Task_4() {
	arr := []int{11, 22}
	fmt.Println(arr)

	n1 := &arr[0]
	n2 := &arr[1]

	arr[0], arr[1] = *n2, *n1

	fmt.Println(arr)
}

// Реализуй функцию, которая принимает указатель на строку и заменяет в ней пробелы на "_".
func Task_5(str *string) {
	r := []rune(*str)

	for i := 0; i < len(r); i++ {
		if string(r[i]) == " " {
			r[i] = '_'
		}
	}

	fmt.Println(string(r))
}

// Создай структуру Person, хранящую имя и возраст, и напиши функцию, изменяющую возраст через указатель.
type Person struct {
	Age  int
	Name string
}

func Task_6(p *Person) {
	fmt.Println(p.Age, p.Name)
	p.Age = 10
	fmt.Println(p.Age, p.Name)
}

// Реализуй функцию, которая принимает указатель на массив из 5 чисел и заменяет все чётные числа на их квадраты.
func Task_7(arr *[5]int) {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			arr[i] = int(math.Pow(float64(arr[i]), 2))
		}
	}
	fmt.Println(arr)
}

// Напиши функцию, которая принимает указатель на map[string]int и увеличивает все значения на 5.
func Task_8(myMap map[string]int) {

	for _, value := range myMap {
		value += 5
	}
	fmt.Println(myMap)

}

// Создай структуру Rectangle (ширина, высота) и реализуй метод Scale, который увеличивает размеры на заданный коэффициент через указатели.
type Rectangle struct {
	Width  float64
	Height float64
}

// func Task_9(r *Rectangle)
func Scale(r *Rectangle, value *float64) {
	fmt.Println(r.Width, r.Height)

	r.Height = *value + float64(r.Height)
	r.Width = *value + float64(r.Width)

	fmt.Println(r.Width, r.Height)
}
