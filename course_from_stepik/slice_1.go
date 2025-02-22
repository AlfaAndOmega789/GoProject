package main

import "fmt"

func main() {
	//добавление элементов
	var buf []int // len= 0, cap= 0
	fmt.Println(buf)
	buf = append(buf, 9, 10) //len= 2, cap= 2
	fmt.Println(buf)
	buf = append(buf, 12) //len= 3, cap=4
	fmt.Println(buf)

	//добавление другого слайса
	otherBuf := make([]int, 3)     //[0,0,0]
	buf = append(buf, otherBuf...) // len=6, cap=8
	fmt.Println(buf, otherBuf)

	var bufLen, bufCap int = len(buf), cap(buf)
	fmt.Print(bufLen, bufCap) //просмотр информации о слайсе
}
