package main

import "fmt"

type myOwnType int

var x myOwnType

func main() {
	x = 10

	fmt.Printf("x: %v, %T\n", x, x)

	y := int(x)

	fmt.Printf("y: %v, %T", y, y)
}
