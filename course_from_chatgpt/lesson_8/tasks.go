package lesson_8

import (
	"fmt"
	"math"
	"reflect"
)

// Создай интерфейс Shape с методами Area() и Perimeter(). Реализуй его для структур Circle и Rectangle.
//type Shape interface {
//	Area()
//	Perimeter()
//}
//
//type Circle struct {
//	Radius float64
//}
//type Rectangle struct {
//	Width  float64
//	Height float64
//}
//
//func (c Circle) Area() {
//
//}
//func (r Rectangle) Perimeter() {
//
//}

// Напиши интерфейс Mover с методом Move(), который реализуют структуры Car и Person.
//type Mover interface {
//	Move()
//}
//type Car struct {
//	Name string
//}
//type Person struct {
//	Name string
//}
//
//func (c Car) Move() {
//	fmt.Println(c.Name)
//}
//func (p Person) Move() {
//	fmt.Println(p.Name)
//}
//
//func Info(m Mover) {
//	m.Move()
//}
//
//func Task_2() {
//	c := Car{Name: "BMW"}
//	p := Person{Name: "Ivan"}
//
//	Info(c)
//	Info(p)
//}

// Создай интерфейс Animal с методом Sound(), реализуй его для Cat и Dog.
//type Animal interface {
//	Sound()
//}
//
//type NewCar struct {
//	NewSound string
//}
//
//type NewDog struct {
//	NewSound string
//}
//
//func (c NewCar) Sound() {
//	fmt.Println(c.NewSound)
//}
//func (d NewDog) Sound() {
//	fmt.Println(d.NewSound)
//}
//
//func Info(a Animal) {
//	a.Sound()
//}
//
//func Task_3() {
//	c := NewCar{"Pi-pi"}
//	d := NewDog{"Gaw"}
//
//	Info(c)
//	Info(d)
//}

// Создай интерфейс Payment с методом Pay(amount float64), реализуй его для CreditCard и PayPal.
//type Payment interface {
//	Pay(amount float64)
//}
//type CreditCard struct {
//}
//type PayPal struct {
//}
//
//func (c CreditCard) Pay(amount float64) { fmt.Println("CreditCard") }
//func (p PayPal) Pay(amount float64)     { fmt.Println("PayPal") }
//
//func NewInfo(p Payment) {
//	p.Pay(123.456)
//}
//
//func Task_4() {
//	c := CreditCard{}
//	p := PayPal{}
//
//	NewInfo(c)
//	NewInfo(p)
//}

// Напиши функцию, которая принимает срез Shape и вычисляет суммарную площадь всех фигур.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * 3.14
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func totalArea(shapes []Shape) float64 {
	var sum float64
	for i := 0; i < len(shapes); i++ {
		if reflect.TypeOf(shapes[i]) == reflect.TypeOf(Circle{}) {
			sum += shapes[i].Area()
		}
		if reflect.TypeOf(shapes[i]) == reflect.TypeOf(Rectangle{}) {
			sum += shapes[i].Area()
		}
	}
	return sum
}

func Task_5() {
	cir := Circle{Radius: 50.5}
	rec := Rectangle{Width: 10.5, Height: 20.0}

	shapes := []Shape{cir, rec}

	fmt.Println(totalArea(shapes))
}
