package main

import "fmt"

func main() {

	puf := []int{1, 2, 3, 4, 5}
	fmt.Println(puf)

	//получение среза указывающую на ту же память
	sl1 := puf[1:4] //[2, 3, 4]
	sl2 := puf[:2]  //[1, 2]
	sl3 := puf[2:]  //[3, 4, 5]
	fmt.Println(sl1, sl2, sl3)

	newPuf := puf[:] //[1, 2, 3, 4, 5]
	fmt.Println(newPuf)
	newPuf[0] = 9 // [9, 2, 3, 4, 5]
	fmt.Println(newPuf)

	// newPuf теперь указывает на другие данные
	newPuf = append(newPuf, 6)

	//puf = [9, 2, 3, 4, 5], не изменился
	//newBuf = [1, 2, 3, 4, 5, 6], изменился
	newPuf[0] = 1
	fmt.Println("puf", puf)
	fmt.Println("newPuf", newPuf)

	//копирование одного слайса в другой
	var emptyPuf []int // len=0 , cap=0
	//неправильно - скопирует меньше (по len) из 2 слайсов
	copied := copy(emptyPuf, puf) // copied 0
	fmt.Println(copied, emptyPuf)

	// правильно
	newPuf = make([]int, len(puf), len(puf))
	copy(newPuf, puf)
	fmt.Println(newPuf)

	//можно копировать в часть существующего слайса
	ints := []int{1, 2, 3, 4}
	//fmt.Print(ints)
	copy(ints[1:3], []int{5, 6}) //ints = [1, 5, 6, 4] вместо 2,3 -> вставили 5, 6
	fmt.Println(ints)
}
