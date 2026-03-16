package main

import (
	// a-tour-of-go is a module/project
	// "a-tour-of-go/basics" // This is a package/folder
	"a-tour-of-go/flow_control_statements"

	"fmt"
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

// Outside a function, every statement begins with a keyword
// (var, func, and so on) and so the := construct is not available.

func main() {
	fmt.Println("Genesis Corp MAIN")

	// Call the HelloGenesisCorp function from the basics package
	// "H" the capital letter means exported
	// basics.Basics()

	// Call the Flows function from the flow_control_statements package
	// flow_control_statements.Flows()
	n := flow_control_statements.DeferStatement()
	fmt.Println(n)
	// flow_control_statements.SqrtExercise()

	// // %v can be used for everything
	// fmt.Printf("My name is: %v %v \n", math.Sqrt(7), "Gen")

	// fmt.Println(add(42, 42))

	// a, b := swap("corp", "genesis")
	// fmt.Println(a, b)

	// b, a = a, b

	// // Wrong usage
	// // fmt.Println(a, b, 123, "sdaads", split(17))

	// fmt.Println(split(17))

	// x, y := split(17)
	// fmt.Println(x, y)

	// // Blank Identifier
	// x2, _ := split(17)
	// fmt.Println(x2)

	// // The var statement declares a list of variables;
	// // as in function argument lists, the type is last.
	// // A var statement can be at package or function level.
	// // var c, python, java bool = true, false, false
	// c, python, java := true, false, false

	// // var i int
	// // Inside a function, the := short assignment statement
	// // can be used in place of a var declaration with implicit type.
	// i := 0
	// fmt.Println(c, python, java, i)

	// var c2, python2, java2 = true, 1453, "no"
	// var k, j int = 1, 2
	// fmt.Println(k, j, c2, python2, java2)

	// // var g = 3
	// test := 3
	// fmt.Println(test)
}
