package main

import (
	"a-tour-of-go/basics"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Genesis Corp MAIN")
	basics.HelloGenesisCorp()

	for {
		var randNum = rand.Intn(3)
		fmt.Println(randNum)
		if randNum == 0 {
			break
		}
	}
}
