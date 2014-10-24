package main

import "fmt"
//MAIN OMIT
func main() {
	strChan := make(chan string)

	for i := 0; i < 3; i++ {
		go World(i, strChan)
	}

	for i := 0; i < 3; i++ {
		go func(index int) {
			strChan <- Hello(index)
		}(i)
	}

	for i := 0; i < 6; i++ {
		fmt.Println(<- strChan)
	}
}
//MAINEND OMIT
//HELLO OMIT
func Hello(index int) string {
	return fmt.Sprintf("Hello from %v", index)
}
//HELLOEND OMIT
//WORLD OMIT
func World(index int, out chan <- string) {
	out <- fmt.Sprintf("World from %v", index)
}
//WORLDEND OMIT