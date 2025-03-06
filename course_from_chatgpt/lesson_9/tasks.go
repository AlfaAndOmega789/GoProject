package lesson_9

import "fmt"

type Writer interface {
	Write()
}

type Pen struct {
	Name string
}

func (p Pen) Write() {
	fmt.Println("Pen: ", p.Name)
}
func Task_1() {
	p1 := Pen{"Pen"}
	var w1 Writer = &p1
	var w2 Writer = p1
	w1.Write()
	w2.Write()
}

type Notifier interface {
	Notify()
}
type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Notify() {
	fmt.Println(e.Email)
}
func SendNotification(n Notifier) {
	n.Notify()
}

func Task_2() {
	e := EmailNotifier{"ivan.koval.andr@gmail.com"}
	var n1 Notifier = e
	var n2 Notifier = &e

	SendNotification(n1)
	SendNotification(n2)
}
