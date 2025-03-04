package chapter_5

import "fmt"

// 1. Объявление и вызов функций (Легкий)
// Напишите функцию greet, которая принимает имя (string) и возвращает строку приветствия. Вызовите эту функцию и выведите результат.
func Greet(name string) {
	fmt.Println("Hi", name)
}

// 2. Имитируем именованные и опциональные параметры (Средний)
// Go не поддерживает опциональные параметры, но их можно имитировать.
// Напишите функцию, которая принимает два числа и опциональный флаг (bool).
// Если флаг true, функция возвращает сумму чисел, иначе разность.
func Imitate(num1 int, num2 int, value bool) int {
	if value {
		return num1 + num2
	}
	return num1 - num2
}

// 3. Вариативные параметры и срезы (Средний)
// Напишите функцию, которая принимает неограниченное количество чисел (variadic parameters) и возвращает их сумму.
func SumValues(vals ...int) int {
	//out := make([]int, 0, len(vals))
	newOut := []int{}
	newOut = append(newOut, vals...)
	var sum int

	//for i := 0; i < len(newOut); i++ {
	//	sum += newOut[i]
	//}

	for _, v := range newOut {
		sum += v
	}

	return sum
}

// 4. Возврат нескольких значений (Средний)
// Создайте функцию, которая принимает два числа и возвращает их сумму и произведение.
func ReturningMultipleValues(num1 int, num2 int) (sum int, multiplication int) {
	sum = num1 + num2
	multiplication = num1 * num2

	return sum, multiplication
}

// 5. Игнорирование возвращаемых значений (Средний)
// Дополните предыдущую задачу: вызовите функцию, но используйте только сумму, игнорируя второе возвращаемое значение с помощью _.
func NewFunction() {
	sum, _ := ReturningMultipleValues(5, 6)
	fmt.Println(sum)
}

// 6. Именованные возвращаемые значения (Средний)
// Реализуйте функцию, которая возвращает два значения, используя именованные возвращаемые параметры.
// Например, функцию, вычисляющую длину и площадь прямоугольника.
func RectagleAreaAndPerimetr(num1 float64, num2 float64) (area float64, perimetr float64) {
	area = num1 * num2
	perimetr = 2 * (num1 * num2)
	return area, perimetr
}

// 7. Анонимные функции (Сложный)
// Создайте и сразу вызовите анонимную функцию, которая принимает два числа и возвращает их разность.
func Anonymous() {
	var value = func(n1 int, n2 int) int { return n1 + n2 }(5, 6)
	fmt.Println(value)
}

// 8. Передача функции в качестве параметра (Сложный)
// Напишите функцию applyOperation, которая принимает два числа и функцию, выполняющую арифметическую операцию (сложение, умножение и т. д.).
// Проверьте, как работает передача функций.
func ApplyOperation(num1 int, num2 int, value func(i int, j int) (int, int)) (sum int, multi int) {
	return value(num1, num2)
}

// 9. Замыкания (Сложный)
// Создайте замыкание, которое увеличивает внутренний счетчик на 1 при каждом вызове и возвращает его текущее значение.
func ExternalMethod() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

// 10. Возвращение функции (Сложный)
// Реализуйте функцию, которая возвращает другую функцию. Например, makeMultiplier(n), которая возвращает функцию, умножающую число на n.
func MakeMultiplier(n int) func() int {
	value := 10
	return func() int { return value * n }
}
