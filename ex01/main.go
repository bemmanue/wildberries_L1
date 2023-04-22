package main

import (
	"fmt"
)

// Human структура, которая хранит поля "возраст" и "имя"
type Human struct {
	age  int
	name string
}

// introduce метод структуры Human, который печатает имя и возраст
func (h *Human) introduce() {
	fmt.Printf("Hello! My name is %s. I am %d years old.", h.name, h.age)
}

// Action структура, которая наследует поля и методы структуры Human
type Action struct {
	Human
}

func main() {
	// Создаем и инициализируем структуру Human
	human := Human{
		age:  18,
		name: "Yulya",
	}

	// Создаем и инициализируем структуру Action
	action := Action{
		Human: human,
	}

	// Вызываем наследуемый от структуры Human метод
	action.introduce()

	// Также получаем доступ к полям, унаследованным от родительской структуры Human
	action.age = 25
	action.name = "Kamil"
}
