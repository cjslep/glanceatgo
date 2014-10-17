package main

import "fmt"
//START OMIT
type metaDataError struct {
	errString string
	errCode int
}

func (m metaDataError) Error() string {
	return m.errString
}

func printError(e error) {
	fmt.Println(e)
}

func main() {
	m := metaDataError{"Flux capacitor requires 1.21 gigawatts", 121}
	printError(m)
	fmt.Println(m.errCode)
}
//END OMIT