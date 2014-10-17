package main

import "fmt"
//START OMIT
func fact(n int) func(int) int { // HL000
	val := 1
	for i := n; i > 0; i-- {
		val *= i
	}
	// Function to access value:
	return func(shift int) int { // HL000
		return val + shift // HL000
	} // HL000
}

func main() {
	fn := fact(3)
	fmt.Println(fn(1))
	fmt.Println(fn(2))
}
//END OMIT