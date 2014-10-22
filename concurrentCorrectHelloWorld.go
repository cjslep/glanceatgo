package main

import "fmt"

func main() {
	//ANONFUNC OMIT
	doneChan := make(chan string) // HL001

	go func() {
		// Do work...
		fmt.Println("Hello World!")
		// Senders close channels
		close(doneChan) // HL001
	}()
	//ANONFUNCEND OMIT

	fmt.Println("Waiting to quit...")
	// Block until we receive something (closed channels return 0 value)
	<-doneChan // HL001
	fmt.Println("Quitting.")
}
