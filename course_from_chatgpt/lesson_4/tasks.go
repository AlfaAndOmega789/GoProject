package lesson_4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Создай map, где ключ — это имя студента, а значение — его оценка. Заполни его 3-4 значениями и выведи их.
func Task_1() {
	myMap := make(map[string]int)

	myMap["Ivan"] = 9
	myMap["Andrei"] = 8
	myMap["Alina"] = 10

	fmt.Println(myMap)
}

// Напиши программу, которая запрашивает у пользователя имя и ищет его в map. Если имя найдено, выведи оценку. Если нет — сообщи об этом.
func Task_2() {
	myMap := make(map[string]int)
	var name string

	myMap["Ivan"] = 9
	myMap["Andrei"] = 8
	myMap["Alina"] = 10

	fmt.Println("Введите имя:")
	fmt.Scanln(&name)

	for key, value := range myMap {
		if key == name {
			fmt.Println("Имя найдено! Оценка студента составляет", value)
			return
		}
	}
	fmt.Println("Имя не найдено!")
}

// Удали из map студента с заданным именем.
func Task_3() {
	myMap := make(map[string]int)

	myMap["Ivan"] = 9
	myMap["Andrei"] = 8
	myMap["Alina"] = 10

	fmt.Println(myMap)

	delete(myMap, "Andrei")

	fmt.Println(myMap)
}

// Перебери все элементы map и выведи их.
func Task_4() {
	myMap := make(map[string]int)

	myMap["Ivan"] = 9
	myMap["Andrei"] = 8
	myMap["Alina"] = 10

	for key, value := range myMap {
		fmt.Println(key, value)
	}
}

// Найди студента с максимальной оценкой.
func Task_5() {
	myMap := make(map[string]int)
	max := 0

	myMap["Max"] = 10
	myMap["Ivan"] = 9
	myMap["Andrei"] = 8
	myMap["Alina"] = 10

	for _, value := range myMap {
		if value > max {
			max = value
		}
	}

	for key, value := range myMap {
		if value == max {
			fmt.Println("Cтудент(ы) с максимальной оценкой:", key)
		}
	}

}

// Создай map, где ключ — это название страны (string), а значение — её столица (string).
// Заполни карту минимум 3 странами и выведи её содержимое.
func Task_6() {
	myMap := make(map[string]string)
	myMap["Belarus"] = "Minsk"
	myMap["Russia"] = "Moscow"
	myMap["Ukraine"] = "Kiev"

	fmt.Println(myMap)
}

// Запроси у пользователя название страны и выведи её столицу (если страна есть в map).
// Если страны нет — выведи сообщение "Страна не найдена".
func Task_7() {
	myMap := make(map[string]string)
	myMap["Belarus"] = "Minsk"
	myMap["Russia"] = "Moscow"
	myMap["Ukraine"] = "Kiev"
	var country string

	fmt.Println("Введите название страны:")
	fmt.Scanln(&country)

	for key, _ := range myMap {
		if key == country {
			fmt.Println(key)
			return
		}
	}
	fmt.Println("Страна не найдена")
}

// Создай map с тремя странами и столицами. Запроси у пользователя название страны и удали её из map. Выведи обновленный map.
func Task_8() {
	myMap := make(map[string]string)
	myMap["Belarus"] = "Minsk"
	myMap["Russia"] = "Moscow"
	myMap["Ukraine"] = "Kiev"
	var country string

	fmt.Println("Введите название страны:")
	fmt.Scanln(&country)

	delete(myMap, country)

	fmt.Println(myMap)

}

//Запроси у пользователя строку, разбей её на слова (используй strings.Fields()),
//создай map, где ключ — слово, а значение — сколько раз оно встречается в строке.
//Выведи результат.

func Task_9() {
	myMap := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку:")
	str, _ := reader.ReadString('\n') //читаем всю строку целиком
	str = strings.TrimSpace(str)      //убираем лишние переводы строк и пробелы

	fields := strings.Fields(str)

	for i := 0; i < len(fields); i++ {
		value := 0
		for j := 0; j < len(fields); j++ {
			if fields[i] == fields[j] {
				value += 1
			}
		}
		myMap[fields[i]] = value
	}

	fmt.Println(myMap)

}

// Создай map, где ключ — имя студента, а значение — его оценка (int).
// Найди и выведи имя студента с максимальной оценкой.
func Task_10() {
	myMap := make(map[string]int)
	max := 0

	myMap["Max"] = 10
	myMap["Ivan"] = 9
	myMap["Andrei"] = 8
	myMap["Alina"] = 10

	for _, value := range myMap {
		if value > max {
			max = value
		}
	}

	for key, value := range myMap {
		if value == max {
			fmt.Println("Cтудент(ы) с максимальной оценкой:", key)
		}
	}
}

// Создай map, где ключ — это оценка студента (int),
// а значение — список имён студентов с такой оценкой ([]string).
// Заполни map данными и выведи её.
func Task_12() {
	myMap := make(map[int][]string)

	myMap[5] = []string{"Ivan", "Max", "Andrei", "Alina"}
	myMap[4] = []string{"Alex", "Olga"}

	fmt.Println(myMap)
}

// Создай две map, содержащие названия товаров и их цены.
// Объедини их в одну map (если товар встречается в обоих map, оставь цену из второй).
func Task_11() {

}
