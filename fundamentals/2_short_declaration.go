package main

import (
	"fmt"
)

var pkgVar = "Package var value" // package scope
var intPkgVar int = 1            // declaring type of var

func main() {
	// The := short variable assignment operator indicates that short variable declaration is being used.
	// There is no need to use the var keyword or declare the variable type.

	// - Auto assign the type of the variable based on the initialized value
	// - MUST declare at least one variable
	// - Block scope

	num := 1
	str := "Good morning!"

	num, floatNum := 2, 2.10

	// Printf args
	// %v -> value of variable
	// %T -> type of variable
	fmt.Printf("num: %v, %T\n", num, num)
	fmt.Printf("str: %v, %T\n", str, str)
	fmt.Printf("floatNum: %v, %T\n", floatNum, floatNum)
	fmt.Printf("pkgVar: %v, %T\n", pkgVar, pkgVar)
	fmt.Printf("pkgVar: %v, %T\n", intPkgVar, intPkgVar)
}
