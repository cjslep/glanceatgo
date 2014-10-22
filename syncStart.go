package main

import (
	"fmt"
	"time"
)
//MAIN OMIT
func main() {
	strChan := make(chan string)
	startChan := make(chan struct{}) // HL

	go World(strChan, startChan) // HL

	go func() {
		fmt.Println("Anon func waiting...")
		<- startChan // HL
		fmt.Println("Anon func executing")
		strChan <- Hello()
	}()

	time.Sleep(100) // Give goroutines a chance to start
	fmt.Println("Main about to send starting signal")
	close(startChan) // Close to let everyone receive 0 value // HL
	fmt.Println("Starting signal sent")
	fmt.Println(<- strChan)
	fmt.Println(<- strChan)
}
//MAINEND OMIT
func Hello() string {
	return "Hello"
}
//WORLD OMIT
func World(out chan <- string, start <- chan struct{}) { // HL
	fmt.Println("World waiting...")
	<- start // Block until we get a value // HL
	fmt.Println("World executing")
	out <- "World"
}
//WORLDEND OMIT