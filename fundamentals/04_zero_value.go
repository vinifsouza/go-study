package main

import (
	"fmt"
)

// Zero is the default value of a var when var is not initialized
var a int
var b float64
var c string
var d bool

func main() {
	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)
	fmt.Printf("d: %v, %T\n", d, d)
}
