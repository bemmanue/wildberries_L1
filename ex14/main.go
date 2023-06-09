package main

import (
	"fmt"
	"reflect"
)

// checkType с помощью пакета reflect получает и возвращает тип переменной
func checkType(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}

func main() {
	var n int
	var s string
	var b bool
	var ch chan int

	fmt.Println(checkType(n))
	fmt.Println(checkType(s))
	fmt.Println(checkType(b))
	fmt.Println(checkType(ch))
}
