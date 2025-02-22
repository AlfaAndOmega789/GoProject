package lesson_2

import "fmt"

func Lesson_2() {

	//Условный оператор if-else
	//if условие {
	//	//код если условие истинно
	//} else if другое_условие {
	//	//код, если второе условие истинно
	//} else {
	//	// код, если ни одно из условий не подошло
	//}
	//
	//var number int
	//fmt.Println("Введите число")
	//fmt.Scanln(&number)
	//
	//if number > 0 {
	//	fmt.Println("Число положительное")
	//} else if number < 0 {
	//	fmt.Println("Число отрицательное")
	//} else {
	//	fmt.Println("Число равно нулю")
	//}

	//switch-case
	//var day int
	//fmt.Println("Введите номер дня недели")
	//fmt.Scanln(&day)
	//
	//switch day {
	//case 1:
	//	fmt.Println("Понедельник")
	//case 2:
	//	fmt.Println("Вторник")
	//case 3:
	//	fmt.Println("Среда")
	//case 4:
	//	fmt.Println("Четверг")
	//case 5:
	//	fmt.Println("Пятница")
	//case 6, 7:
	//	fmt.Println("Выходной")
	//default:
	//	fmt.Println("Некорректное число")
	//}

	//The only loop in Go is for
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("Этот код будет выполняться вечно!")
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i == 7 {
			break
		}
		fmt.Println(i)
	}
}
