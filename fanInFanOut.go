package main

import (
	"fmt"
	"time"
)
//MAIN OMIT
func main() {
	strChan := make(chan string)
	startChan := make(chan struct{}) // HL

	for i := 0; i < 3; i++ {
		go World(i, strChan, startChan) // HL
	}

	for i := 0; i < 3; i++ {
		go func(index int) {
			<- startChan // HL
			strChan <- Hello(index)
		}(i)
	}

	time.Sleep(100) // Give goroutines a chance to start
	// Close to let everyone receive 0 value // HL
	close(startChan) // HL
	for i := 0; i < 6; i++ {
		fmt.Println(<- strChan)
	}
}
//MAINEND OMIT
func Hello(index int) string {
	return fmt.Sprintf("Hello from %v", index)
}
//WORLD OMIT
func World(index int, out chan <- string, start <- chan struct{}) { // HL
	<- start // Block until we get a value // HL
	out <- fmt.Sprintf("World from %v", index)
}
//WORLDEND OMIT