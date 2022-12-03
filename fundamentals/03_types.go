package main

// Go is a static type language

// var [name] [type]
// type -> when the type is not informed, value auto assign the type of the initialized value
// when the value is not assigned, it is only possible to assign inside a block
var x int = 1 // type explicit and initialized
var y string  // missing initialization, MUST explicit type

func main() {
	y = "Assign value to var y"
}
