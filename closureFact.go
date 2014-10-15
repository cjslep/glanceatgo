package main

import "fmt"
//START OMIT
func fact(n int) (func() int) { // HL
	val := 1
	for i := n; i > 0; i-- {
		val *= i
	}
	// Function to access value:
	return func() int { // HL
		return val // HL
	} // HL
}
//END OMIT

func main() {
	fn := fact(0)
	fmt.Println(fn())
	fn = fact(2)
	fmt.Println(fn())
	fn = fact(3)
	fmt.Println(fn())
	fn = fact(5)
	fmt.Println(fn())
}