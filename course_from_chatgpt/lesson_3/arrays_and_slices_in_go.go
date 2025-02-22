package lesson_3

import "fmt"

func Lesson_3() {
	//объявление массива
	var arr1 [5]int
	fmt.Print(arr1)

	//инициализация с значениями
	numbers := [4]int{1, 2, 3, 4}
	fmt.Println(numbers)

	//Определение размера массива автоматически
	letters := [...]string{"A", "B", "C"}
	fmt.Println(letters)

	//Доступ к элементам массива
	arr := [3]int{10, 20, 30}
	fmt.Println(arr[0])
	arr[1] = 50
	fmt.Println(arr)
	fmt.Println(len(arr))

	//Объявление среза
	var slice1 []int // срез без начального размера
	fmt.Println(slice1)

	//Создание среза на основе массива
	arrNew := [5]int{10, 20, 30, 40, 50}
	slice := arrNew[1:4] // срез с 1 по 4 элемент(не включительно)
	fmt.Println(slice)   //[20, 30, 40]

	//Функция make для создания среза
	new_slice := make([]int, 3) //срез из 3 элементов, все = 0
	fmt.Println(new_slice)

	//Добавление и удаление элементов в срезах
	nums := []int{1, 2, 3}
	nums = append(nums, 4, 5)
	fmt.Println(nums) //[1, 2, 3, 4, 5]

	//Объединение срезов
	a := []int{1, 2}
	b := []int{3, 4}
	a = append(a, b...) // добавляем все элементы 'b' в 'a'
	fmt.Println(a)      //[1, 2, 3, 4]

	//итерации по массиву и срезу(for + range)
	arr_new := [4]string{"cherry", "banana", "apple", "date"}

	for i, v := range arr_new {
		fmt.Println("Индекс:", i, "Значение:", v)
	}

	//копирование срезов(copy)
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println(dst) // [1 2 3]
}
