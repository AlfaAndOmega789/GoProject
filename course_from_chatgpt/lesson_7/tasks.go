package lesson_7

import "fmt"

// Создай структуру Book, хранящую название книги и автора. Создай экземпляр и выведи его в консоль.
func Task_1() {
	b := Book{"Outsider", "Stephen King"}
	fmt.Println(b)
}

type Book struct {
	Name   string
	Author string
}

// Добавь метод PrintInfo() для структуры Book, который будет выводить информацию о книге.
func Printinfo(b Book) {
	fmt.Println("Название книги:", b.Name, "Автор:", b.Author)
}

func Task_2() {
	b := Book{"Outsider", "Stephen King"}
	Printinfo(b)
}

// Создай структуру Car с полями Brand и Year. Напиши функцию, которая принимает указатель на Car и обновляет год выпуска.
type Car struct {
	Brand string
	Year  int
}

func newYearOfRelease(car *Car, newYearOfRelease int) {
	fmt.Println(car.Brand, car.Year)
	car.Year = newYearOfRelease

	fmt.Println(car.Brand, car.Year)

}
func Task_3() {
	car := Car{"BMW", 1998}

	newYearOfRelease(&car, 2025)
}

// Создай структуру Student с полями Name, Age и Grades (срез оценок). Добавь метод, который вычисляет средний балл.
func Task_4() {
	student := Student{"Ivan", 18, []float64{9, 10, 8, 7}}
	fmt.Println("Средний балл:", AverageScore(&student))
}

type Student struct {
	Name   string
	Age    int
	Grades []float64
}

func AverageScore(student *Student) float64 {
	arr := student.Grades
	var sum float64 = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum / float64(len(arr))
}

// Напиши функцию, которая принимает указатель на Student и добавляет новую оценку.
func NewGrade(student *Student, newGrade float64) {
	mergeGrade := append(student.Grades, []float64{newGrade}...)
	fmt.Println(mergeGrade)
}

func Task_5() {
	student := Student{"Ivan", 18, []float64{9, 10, 8, 7}}
	fmt.Println(student.Grades)
	NewGrade(&student, 10)
}

// Создай структуру Rectangle с методами Area() и Perimeter()
type Rectangle struct {
	Wight  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Wight * r.Height
}
func (r *Rectangle) Perimeter() float64 {
	return (r.Wight + r.Height) * 2
}

func Task_6() {
	r := Rectangle{14.5, 15.4}
	fmt.Println(r.Area(), r.Perimeter())
}

// Создай структуру BankAccount, которая хранит Balance. Добавь метод Deposit(amount float64), который пополняет баланс.
func Task_7() {
	b := BankAccount{10500.100500}
	fmt.Println(b.Withdraw(10500.100500))
}

type BankAccount struct {
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

// Task_8()
// Напиши метод Withdraw(amount float64), который снимает деньги, но не позволяет уйти в минус.
func (b *BankAccount) Withdraw(amount float64) float64 {
	if b.Balance-amount >= 0 {
		b.Balance -= amount
	}
	return b.Balance
}

// Реализуй структуру Company, где есть поле Employees (срез структур Employee),
// а у Employee есть Name и Salary. Добавь метод, который вычисляет среднюю зарплату сотрудников.
func Task_9() {
	e1 := Employee{"Ivan", 4000}
	e2 := Employee{"Andrei", 2000}
	c := Company{
		Employees: []Employee{e1, e2},
	}

	fmt.Println(c.averageSalary())

}

type Company struct {
	Employees []Employee
}

type Employee struct {
	Name   string
	Salary float64
}

func (c *Company) averageSalary() float64 {
	var averageSalary float64
	for i := 0; i < len(c.Employees); i++ {
		averageSalary += c.Employees[i].Salary
	}
	if len(c.Employees) == 0 {
		return 0
	}
	return averageSalary / float64(len(c.Employees))
}
