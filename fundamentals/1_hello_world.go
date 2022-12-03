package main

import (
	"fmt"
)

func main() {
	nBytes, errors := fmt.Println("Hello, world!", 100)

	fmt.Println(nBytes, errors)
}
