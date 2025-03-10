package chapter_4

import "fmt"

// 1. Затенение переменных (Легкий)
// Напишите программу, в которой объявляется глобальная переменная,
// затем внутри функции создаётся локальная переменная с таким же именем. Выведите обе переменные, объясните результат.
const value = 1000

func VariableShading() {
	value := 100

	func() {
		fmt.Println(value)
		value := 50
		fmt.Println(value)
	}()
	fmt.Println(value)
}

//2. Выявление затенённых переменных (Легкий)
//Напишите код, в котором объявляется переменная в одном блоке, а затем создаётся переменная с тем же именем во вложенном блоке. Используйте go vet, чтобы выявить проблему.

// 3. Оператор if (Легкий)
// Напишите программу, которая запрашивает у пользователя число и выводит, является ли оно положительным, отрицательным или нулём, используя оператор if.
func SomeMethod() {
	var number int
	fmt.Println("Enter whole numer:")
	fmt.Scanln(&number)

	if number > 0 {
		fmt.Println("positive")
	} else if number < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}
}

//4. Полный оператор for (Средний)
//Используйте for с тремя выражениями (инициализация; условие; изменение), чтобы вывести на экран числа от 10 до 1 в обратном порядке.

func FullFor() {
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}

// 5. for, использующий только условие (Средний)
// Напишите цикл, который уменьшает значение переменной x на 2, пока оно не станет меньше или равно 0. Используйте for, указывая только условие.
func ConditionFor() {
	var x int = 100

	for x >= 0 {
		x -= 2
		fmt.Println(x)
	}
}

//6. Бесконечный for и break (Средний)
//Создайте бесконечный цикл for, который увеличивает число на 1, пока не достигнет 100. Как только число станет 50, используйте break, чтобы прервать выполнение.

func InfiniteFor() {
	var number int
	for number < 100 {
		number++

		if number == 50 {
			break
		}
	}
}

//7. for-range (Средний)
//Создайте массив из 5 чисел. Используйте for-range, чтобы найти сумму всех элементов массива.

func SumArray() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	var sum int

	for _, v := range array {
		sum += v
	}
	fmt.Println(sum)
}

// 8. Оператор switch (Средний)
// Напишите программу, которая запрашивает у пользователя день недели и выводит, является ли этот день будним или выходным, используя switch.
func DayOfTheWeek() {
	var day int
	fmt.Println("Введите день недели(1-7):")
	fmt.Scanln(&day)

	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Вы ввели что-то не то")
	}
}

//9. if vs. switch (Сложный)
//Реализуйте одну и ту же логику (определение сезона по месяцу) двумя способами: через if-else и через switch.
//Сравните читаемость кода и объясните, какой вариант предпочтительнее.
