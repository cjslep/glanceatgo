package main

import "fmt"
//MAIN OMIT
func main() {
	strChan := make(chan string)

	for i := 0; i < 3; i++ {
		go World(i, strChan)
	}

	for i := 0; i < 3; i++ {
		go func(index int) { // HL
			strChan <- Hello(index)
		}(i) // HL
	}

	for i := 0; i < 6; i++ {
		fmt.Println(<- strChan)
	}
}
//MAINEND OMIT
func Hello(index int) string {
	return fmt.Sprintf("Hello from %v", index)
}

func World(index int, out chan <- string) {
	out <- fmt.Sprintf("World from %v", index)
}