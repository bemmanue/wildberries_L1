package main

import (
	"fmt"
	"math/big"
)

func main() {
	//  создаем переменные, которые могут хранить большие целочисленные значения
	a := new(big.Int)
	b := new(big.Int)
	c := new(big.Int)

	// присваиваем значения
	a.SetString("3456789087654376565784356789098765", 10)
	b.SetString("108976543354657890876543546789", 10)

	fmt.Printf("a = %d\nb = %d\n\n", a, b)

	// производим операции
	fmt.Printf("a + b = %d\n", c.Add(a, b))
	fmt.Printf("a - b = %d\n", c.Sub(a, b))
	fmt.Printf("a * b = %d\n", c.Mul(a, b))
	fmt.Printf("a / b = %d\n", c.Div(a, b))

}
