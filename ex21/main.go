package main

import "fmt"

// Target описывает требуемый интерфейс
type Target interface {
	Greet()
}

// Adaptee – струтура, которую необходимо адаптировать под интерфейс
type Adaptee struct {
}

// SpecificGreet метод структуры Adaptee
func (a *Adaptee) SpecificGreet() {
	fmt.Println("Hello")
}

// Adapter адаптер, реализующий интерфейс Target
type Adapter struct {
	*Adaptee
}

// Greet метод адаптера
func (a *Adapter) Greet() {
	a.SpecificGreet()
}

// NewAdapter конструктор адаптера
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

// функция, принимающая в виде аргумент в виде объекта, реализующего интерфейс Target
func check(target Target) {
	target.Greet()
}

func main() {
	var adaptee Adaptee

	// с помощью адаптера преобразуем Adaptee в необходимый объект интерфейса Target
	target := NewAdapter(&adaptee)

	check(target)
}
