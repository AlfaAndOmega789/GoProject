package chapter_3

import "fmt"

// 1
func NewSlice() (int, int) {
	data := []int{2, 4, 6, 8, 10}

	return cap(data), len(data)
}

// 2
func Slice() {
	var data []int

	for i := 1; i <= 10; i++ {
		data = append(data, i)
	}

	fmt.Println(data)
}

// 3
func MakeSlice() {
	var data = make([]int, 5, 10)

	fmt.Println(data)
}

// 4
func ElementsSlice() {
	var data []int = []int{1, 2, 3, 4, 5}

	var new1 []int = data[:3]
	var new2 []int = data[len(data)-2 : len(data)]

	fmt.Println(data)
	fmt.Println(new1, new2)
}

func ArrayToSlice1() {
	var array [5]int = [5]int{10, 20, 30, 40, 50}
	var slice []int = array[:]

	fmt.Println(array)
	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		slice[i] += 1
	}

	fmt.Println(array)
	fmt.Println(slice)
}

// 5 Почему именно так?
func ArrayToSlice2() {
	var array [5]int = [5]int{10, 20, 30, 40, 50}
	var slice []int = array[:]

	fmt.Println(array)
	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		slice[i] += 1
	}

	fmt.Println(array)
	fmt.Println(slice)
}

// 6
func CopySlice() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{10, 11, 12}

	newSlice1 := copy(slice1, slice2)
	newSlice2 := copy(slice2, slice1)

	fmt.Println(slice1, newSlice1)
	fmt.Println(slice2, newSlice2)
}

// 7
func DeleteElementInSlice() {
	//удалить элемент без использования append, например 2 элемент
	var number int = 2
	var str string
	slice1 := []string{"a", "b", "c", "d"}

	fmt.Println(str)
	for k, v := range slice1 {
		if k != number {
			str += v
		}
	}
	fmt.Println(str)
}

// 8
func removeDuplicates(s []int) []int {
	newMap := make(map[int]int)
	var slice []int = []int{}

	for k, v := range s {
		newMap[v] = k
	}

	for k, _ := range newMap {
		slice = append(slice, k)
	}

	return slice
}

// 9
func rotateLeft(s []int, n int) {

	for i := 0; i < n; i++ {
		s = OneRotateLeft(s)
	}

	fmt.Println(s)
}

func OneRotateLeft(s []int) (newSlice []int) {
	slice := s[1:]
	firtstElement := s[0]

	newSlice = append(slice, firtstElement)
	return newSlice
}

// 10
func Push(slice []int, n int) []int {
	return append(slice, n)
}

func Pop(slice []int) []int {
	return slice[:len(slice)-1]
}
func Top(slice []int) int {
	return slice[len(slice)-1]
}
