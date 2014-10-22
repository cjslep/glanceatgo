package main

import (
	"fmt"
	"time"
)
//CHAN OMIT
func main() {
	helloChan := make(chan string) // HL
	worldChan := make(chan string) // HL
	startChan := make(chan struct{})

	for i := 0; i < 50; i++ {
		go World(i, worldChan, startChan) // HL
	}

	for i := 0; i < 50; i++ {
		go func(index int) {
			<- startChan
			helloChan <- Hello(index) // HL
		}(i)
	}
	//CHANEND OMIT
	time.Sleep(100)
	//MAIN OMIT
	close(startChan)
	for i := 0; i < 100 ; i++ {
		select {
		case w := <- worldChan: // HL
			fmt.Println("Hello " + w)
		case h := <- helloChan: // HL
			fmt.Println("You say goodbye, and I say " + h)
		}
	}
	//MAINEND OMIT
}
func Hello(index int) string {
	return fmt.Sprintf("Hello (from %v)", index)
}
//WORLD OMIT
func World(index int, out chan <- string, start <- chan struct{}) { // HL
	<- start
	out <- fmt.Sprintf("World from %v", index)
}
//WORLDEND OMIT