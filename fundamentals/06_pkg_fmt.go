package main

import (
	"fmt"
)

func main() {
	fmt.Println("Println") // add new line at end

	x := "Hi"
	y := "Good morning"
	z := fmt.Sprint(x, " ", y) // assign string to a var and non display value

	fmt.Println(z)
}
