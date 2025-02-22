package lesson_3

import "fmt"

// Создай массив из 5 целых чисел и выведи его.
func Task_1() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
}

// Объяви срез строк и добавь в него три города, затем выведи их.
func Task_2() {
	arr_str := []string{}
	arr_str = append(arr_str, "Minsk", "Moscow", "Brest")
	fmt.Println(arr_str)
}

// Запроси у пользователя 5 чисел, сохрани их в срезе и выведи сумму.
func Task_3() {
	arr := []int{}
	var num int
	var sum int
	fmt.Println("Введите 5 целых чисел.")

	for i := 1; i <= 5; i++ {
		fmt.Println(i, ":")
		fmt.Scanln(&num)
		arr = append(arr, num)
		sum += num
	}

	fmt.Println(arr)
	fmt.Println(sum)
}

// Выведи все элементы среза в обратном порядке.
func Task_4() {
	arr := []int{1, 2, 3, 4, 5}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}
}

// Найди минимальное и максимальное значение в срезе.
func Task_5() {
	arr := []int{1, 3, 5, 7, 9}
	min_value, max_value := arr[0], arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < min_value {
			min_value = arr[i]
		}
		if arr[i] > max_value {
			max_value = arr[i]
		}
	}
	fmt.Println(min_value, max_value)
}

// Удали из среза элемент с индексом k, введенным пользователем.
func Task_6() {
	slice := []int{1, 2, 3}
	fmt.Println("Введите индекс элемента требуемый для удаления:")
	var a int
	fmt.Scanln(&a)

	slice = append(slice[:a], slice[a+1:]...)
	fmt.Println(slice)

}
