package main

import "fmt"
//MAIN OMIT
func couldReturnEarly(b bool) {
	fmt.Println("Before defer.")

	defer func() { // HL
		fmt.Printf("I received: %v\n", b) // HL
	}() // HL

	fmt.Println("After defer.")
	if b {
		fmt.Println("Early return.")
		return 
	}
	fmt.Println("Late return.")
	return
}

func main() {
	couldReturnEarly(true)
	fmt.Println()
	couldReturnEarly(false)
}
//MAINEND OMIT