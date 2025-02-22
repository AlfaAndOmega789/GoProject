package lesson_4

import "fmt"

func Lesson_4() {
	var myMap map[string]int     // nil map (нельзя использовать без инициализации)
	myMap = make(map[string]int) // создаем пустую map
	fmt.Println(myMap)

	myNewMap := make(map[string]int)
	fmt.Println(myNewMap)

	//Операции с map
	//Добавление и изменение элементов
	myMap1 := make(map[string]int)
	myMap1["apple"] = 5
	myMap1["banana"] = 7
	myMap1["apple"] = 10

	//Чтение значений
	fmt.Println(myMap1["apple"])

	//Проверка наличия ключа
	value, exists := myMap1["banana"]
	if exists {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Ключ отсутствует")
	}

	//Удаление элемента
	delete(myMap1, "banana") //удалит "banana"

	//Перебор элементов(for range)
	for key, value := range myMap1 {
		fmt.Println("Ключ:", key, "Значение:", value)
	}
}
