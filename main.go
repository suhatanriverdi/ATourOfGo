package main

import (
	// a-tour-of-go is a module/project
	"a-tour-of-go/basics" // This is a package/folder
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Genesis Corp MAIN")

	// Call the HelloGenesisCorp function from the basics package
	basics.HelloGenesisCorp()

	for {
		var randNum = rand.Intn(3)
		fmt.Println(randNum)
		if randNum == 0 {
			break
		}
	}

	fmt.Printf("My name is: %g \n", math.Sqrt(7))
}
