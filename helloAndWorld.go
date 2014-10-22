package main

import "fmt"
//MAIN OMIT
//BLOCK OMIT
func main() {
	strChan := make(chan string) // HL001

	for i := 0; i < 3; i++ {
		go World(i, strChan)
	}

	for i := 0; i < 3; i++ {
		strChan <- Hello(i) // HL001
	}
	//BLOCKENDOMIT

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