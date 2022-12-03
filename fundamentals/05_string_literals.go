package main

import (
	"fmt"
)

func main() {
	// interpreted string literals
	x := "Hello, good morning\nHow are you?\tI hope everything is fine!"

	// raw string literals
	y := `Hello, good morning
How are you?	I hope everything is fine!`

	fmt.Println(x)
	fmt.Println()
	fmt.Println(y)
}
