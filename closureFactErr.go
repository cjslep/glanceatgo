package main

import (
	"fmt"
	"errors"
)

func fact(n int) (int, error) { // HL
	if n < 0 {
		return 0, errors.New("Factorial can only handle nonnegative integers") // HL
	}
	val := 1
	for i := n; i > 0; i-- {
		val *= i
	}
	return val, nil // HL
}

func main() {
	fmt.Println(fact(0))
	fmt.Println(fact(-1))
	fmt.Println(fact(5))
}