package main

import (
	// a-tour-of-go is a module/project
	"a-tour-of-go/basics" // This is a package/folder
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 2
	y = sum - x
	// Naked return
	return
}

func main() {
	fmt.Println("Genesis Corp MAIN")

	// Call the HelloGenesisCorp function from the basics package
	// "H" the capital letter means exported
	basics.HelloGenesisCorp()

	// %v can be used for everything
	fmt.Printf("My name is: %v %v \n", math.Sqrt(7), "Gen")

	fmt.Println(add(42, 42))

	a, b := swap("corp", "genesis")
	fmt.Println(a, b)

	b, a = a, b

	// Wrong usage
	// fmt.Println(a, b, 123, "sdaads", split(17))

	fmt.Println(split(17))

	x, y := split(17)
	fmt.Println(x, y)
	
	// Blank Identifier
	x2, _ := split(17)
	fmt.Println(x2)
}
