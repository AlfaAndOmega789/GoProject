package lesson_5

import (
	"fmt"
	"strings"
)

func Tasks() {
	//printMessage("Ivan")
	//fmt.Println(square(5))
	//fmt.Println(max(6, 7))
	//fmt.Println(factorial(100))
	//fmt.Println(isEven(4))
	//fmt.Println(reverseString("Hello!"))
	//fmt.Println(fibonacci(100))
	//fmt.Println(countVowels("Hello, Ivan!"))

	//arr := []int{5, 5, 6, 1, 5, 3, 0, -1, 7, 0}
	//fmt.Println(findMixMax(arr))
}

// Напиши функцию printMessage(message string), которая принимает строку и выводит её.
func printMessage(str string) {
	fmt.Println(str)
}

// Напиши функцию square(n int) int, которая возвращает квадрат числа.
func square(n int) int {
	return n * n
}

// Напиши функцию max(a, b int) int, которая возвращает большее из двух чисел.
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Напиши функцию factorial(n int) int, которая вычисляет факториал числа.
func factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}

	return fact
}

// Напиши функцию isEven(n int) bool, которая возвращает true, если число чётное, иначе false.
func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

// Напиши функцию reverseString(s string) string, которая переворачивает строку
func reverseString(s string) string {
	runes := []rune(s) //преобразуем строку в руны

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}

// Напиши функцию fibonacci(n int) int, которая возвращает n-е число Фибоначчи.
func fibonacci(n int) int {
	if 2 < n {
		return n
	}
	slice := make([]int, n+1)

	slice[0], slice[1] = 0, 1

	for i := 2; i < 1000; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice[n]
}

// Напиши функцию countVowels(s string) int, которая считает количество гласных в строке.
func countVowels(s string) int {
	s = strings.ToLower(s)
	arr := []rune{'e', 'u', 'i', 'o', 'a', 'y'}
	s_arr := []rune(s)
	count := 0

	for i := 0; i < len(s_arr); i++ {
		for j := 0; j < len(arr); j++ {
			if s_arr[i] == arr[j] {
				count += 1
			}
		}
	}
	return count
}

// Напиши функцию findMinMax(arr []int) (int, int), которая возвращает минимальный и максимальный элемент в срезе.
func findMixMax(arr []int) (min int, max int) {
	min, max = arr[0], arr[0]

	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}
	return min, max
}
