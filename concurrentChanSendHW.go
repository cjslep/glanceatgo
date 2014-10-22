package main

import "fmt"

func main() {
	//ANONFUNC OMIT
	doneChan := make(chan string) // HL001

	go func() {
		// Share memory by communicating
		doneChan <- "Hello World!" // HL001
	}()
	//ANONFUNCEND OMIT

	fmt.Println(<-doneChan) // HL001
}
