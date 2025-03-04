package lesson_11

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Print("Greetings from Goroutine!")
}

func Lesson_11() {
	go sayHello()
	fmt.Println("Greetings from main!")

	time.Sleep(time.Second)

	ch := make(chan int) // Канал для передачи int

	ch <- 5 // Отправить в канал
	//x := <-ch // Получить из канала

}
