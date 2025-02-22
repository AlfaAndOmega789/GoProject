package lesson_2

import "fmt"

// Запроси у пользователя число и выведи "Четное" или "Нечетное" с использованием if-else.
func Task_1() {
	var number int
	fmt.Println("Введите целое число:")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
}

// Запроси у пользователя число от 1 до 12 и выведи соответствующий месяц с помощью switch.
func Task_2() {
	var number int
	fmt.Println("Введите целое число от 1 до 12")
	fmt.Scanln(&number)

	switch number {
	case 1:
		fmt.Println("Январь")
	case 2:
		fmt.Println("Февраль")
	case 3:
		fmt.Println("Март")
	case 4:
		fmt.Println("Апрель")
	case 5:
		fmt.Println("Май")
	case 6:
		fmt.Println("Июнь")
	case 7:
		fmt.Println("Июль")
	case 8:
		fmt.Println("Август")
	case 9:
		fmt.Println("Сентябрь")
	case 10:
		fmt.Println("Октябрь")
	case 11:
		fmt.Println("Ноябрь")
	case 12:
		fmt.Println("Декабрь")
	default:
		fmt.Println("Вы ввели число не входящее в диапазон!")
	}
}

// Выведи все числа от 1 до 10 с использованием for.
func Task_3() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// Запроси у пользователя два числа и выведи все числа между ними.
func Task_4() {
	var firstNumber, secondNumber int

	fmt.Println("Введите первое число:")
	fmt.Scanln(&firstNumber)

	fmt.Println("Введите второе число:")
	fmt.Scanln(&secondNumber)

	if firstNumber == secondNumber {
		fmt.Println("Диапазон не найден")
		return
	}

	if firstNumber > secondNumber {
		firstNumber, secondNumber = secondNumber, firstNumber
	}

	for i := firstNumber; firstNumber < secondNumber; i++ {
		fmt.Println(i)
	}
	//if firstNumber > secondNumber {
	//	for i := secondNumber + 1; i < firstNumber; i++ {
	//		fmt.Println(i)
	//	}
	//} else if secondNumber > firstNumber {
	//	for i := firstNumber + 1; i < secondNumber; i++ {
	//		fmt.Println(i)
	//	}
	//} else {
	//	fmt.Println("Диапазон не найден")
	//}
}

// Используя for, выведи сумму всех чисел от 1 до n, где n вводит пользователь.
func Task_5() {
	var number int
	fmt.Println("Введите число: ")
	fmt.Scanln(&number)
	sum := 0

	for i := 1; i <= number; i++ {
		sum += i
	}

	fmt.Println(sum)
}

// Выведи таблицу умножения от 1 до 10 с помощью вложенных циклов.
func Task_6() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print(i, "*", j, "=", i*j)
			fmt.Println()
		}
		fmt.Println()
	}
}

// Напиши программу, которая ищет все числа от 1 до 100, делящиеся на 3 и 5 одновременно.
func Task_7() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i)
		}
	}

}
