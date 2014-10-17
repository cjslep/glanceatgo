package main

import "fmt"

//START OMIT
func fact(n int) int {
	var val int = 1
	for i := n; i > 0; i-- {
		val *= i
	}
	return val
}
//END OMIT

func main() {
	fmt.Println(fact(0))
	fmt.Println(fact(2))
	fmt.Println(fact(3))
	fmt.Println(fact(5))
}