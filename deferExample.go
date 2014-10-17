package main

import "fmt"

func couldReturnEarly(b bool) {
	fmt.Println("Before defer.")
	defer func() {
		fmt.Printf("I received: %v\n", b)
	}()
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
