package lesson_5

import "fmt"

func sayHello() {
	fmt.Println("Hello, world!")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func sum(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func divide(a, b int) (result int, remainder int) {
	result = a / b
	remainder = a % b
	return
}

func Lesson_5() {
	sayHello()
	greet("Ivan")
	fmt.Println(sum(6, 7))
	fmt.Println(swap(6, 7))
	fmt.Println(divide(10, 20))

}
