package lesson_9

import "fmt"

type Printer interface {
	Print()
}
type Document struct {
	Text string
}

// метод реализуется для значений структуры
func (d Document) Print() {
	fmt.Println("Document: ", d.Text)
}

// метод объявлен через указатель
func (d *Document) NewPrint() {
	fmt.Println("Document: ", d.Text)
}

func Lesson_9() {
	doc := Document{"Go - best programming language!"}
	var p Printer = doc
	p.Print()

	var p2 Printer = &doc
	p2.Print()

	docNew := Document{"Study Go!"}
	//var p4 Printer = docNew
	var p3 Printer = &docNew //ошибка! Document не реализует Printer для значений
	p3.Print()

	s := Service{"Payment"}
	//PrintLog(s)//ошибка реализуется только через указатель
	PrintLog(&s)
}

type Logger interface {
	Log()
}
type Service struct {
	Name string
}

func (s *Service) Log() {
	fmt.Println("Service", s.Name)
}

func PrintLog(l Logger) {
	l.Log()
}
