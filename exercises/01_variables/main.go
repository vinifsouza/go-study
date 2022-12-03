package main

import "fmt"

/*
- Using short declaration, assign this values to the identifiers "x", "y", e "z".
    1. 42
    2. "James Bond"
    3. true
- Display values with:
    1. One print declaration
    2. Multiples print declarations
*/

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("x: %v, y: %v, z: %v\n", x, y, z)
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)
}
