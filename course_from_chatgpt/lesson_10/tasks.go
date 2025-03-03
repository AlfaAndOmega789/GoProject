package lesson_10

import "fmt"

type Address struct {
	Street string
	City   string
	Index  int
}

type User struct {
	Username string
	Email    string
	Address
}

// Создай структуру Address (улица, город, индекс) и встрой её в User.
// Поля Address должны быть доступны в User напрямую.
func Task_1() {
	u := User{
		Username: "kvl.ivan",
		Email:    "ivan.koval.andr@gmail.com",
		Address:  Address{Street: "Lenina", City: "Minsk", Index: 222555},
	}

	fmt.Println(u.Street, u.City, u.Index)
	fmt.Println(u.Address)
}

type Car struct {
	Brand string
}
type ElectricCar struct {
	Car
	BatteryCapacity float64
}

func (c Car) Drive() {
	fmt.Println("The car is moving", c.Brand)
}

func (e ElectricCar) Drive() {
	fmt.Println("The electric car is moving", e.Car.Brand)
}

// Создай структуру Car с полем Brand, а потом структуру ElectricCar, встроив в неё Car.
// Добавь поле BatteryCapacity, реализуй метод Drive() для Car, но переопредели его в ElectricCar.
func Task_2() {
	c := Car{
		Brand: "BMW",
	}
	e := ElectricCar{
		Car:             Car{Brand: "Tesla"},
		BatteryCapacity: 100500.100500,
	}
	c.Drive()
	e.Drive()
}

type Animal struct{}

type Dog struct {
	Animal
}

func (a Animal) MakeSound() {
	fmt.Println("The animal makes a sound")
}
func (d Dog) Bark() {
	fmt.Println("Gaw-gaw!")
}

// Создай структуру Animal с методом MakeSound(), а затем Dog, которая наследует Animal.
// Добавь новый метод Bark() в Dog и вызови оба метода.
func Task_3() {
	d := Dog{}
	d.MakeSound()
	d.Bark()

}
