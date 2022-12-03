package main

import "fmt"

type myOwnType int

var a myOwnType

func main() {
	a = 10

	fmt.Printf("a: %v %T", a, a)
}
