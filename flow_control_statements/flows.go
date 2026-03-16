package flow_control_statements

import (
	"fmt"
	"runtime"
	"time"
)

func printFunc(toBePrinted string, x int) {
	fmt.Println(toBePrinted, x)
	fmt.Println()
}

func testDefer() {
	fmt.Println("Counting...")

	for i := range 10 {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}

/*
* We cna use defer in the following scenarios:
* defer db.Close()
* defer file.Close()
*
* ...
* x := 10
* defer fmt.Println("x's value:", x) // Burada x'i 10 olarak not eder
* x = 20
* ...
* Will print: "x's value: 10", not 20!
 */
func DeferStatement() int {
	var constantine = 10
	// LIFO, stack order here, .. 1 will be printed first
	// Due to the stack ordering, 3 will be put first onto the stack
	defer printFunc("Right before return 3 :)", constantine)
	defer printFunc("Right before return 2 :)", constantine)

	constantine = 1452
	defer printFunc("Right before return 1 :)", constantine)

	fmt.Println("DeferStatement is ended.")

	defer func() {
		// Important Note
		// This is local variable, this won't effect the actual variable
		// inside the DeferStatement() function though
		// This is pass by copied to the stack
		constantine++
	}()

	testDefer()

	// All the above defer's will be executed first,
	// Right before the return here, after all is executed
	// function will then return 1453 finally
	return constantine
}

func Flows() {
	fmt.Println("flow_control_statements package")
	fmt.Println()

	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
		sum += i
	}

	fmt.Println()
	fmt.Println()

	// Alternative for loop
	// This transformation is applied only if
	// (a) the loop variable is not modified within the loop body
	// (b) and the loop's limit expression is not modified within the loop, as `for range` evaluates its operand only once.
	for nth := range 10 {
		fmt.Print(nth, " ")
	}

	// Infinite loop
	// for true {
	// 	fmt.Println()
	// }

	// Infinite loop
	// If you omit the loop condition it loops forever,
	// so an infinite loop is compactly expressed.
	//
	// for {
	// }

	// The init and post statements are optional.
	// Like a while loop
	fmt.Println()

	i := 4
	for i < 7 {
		fmt.Println(i, " ")
		i += 1
	}

	if i == 5 && i < 0 {
		fmt.Println("i is good")
	}

	// the if statement can start with a short statement
	// to execute before the condition.
	if in_var := 1453; in_var == 1453 {
		fmt.Println("Conquest")
	}

	//S witch
	// A switch statement is a shorter way to write a sequence of if - else statements.
	// It runs the first case whose value is equal to the condition expression.
	// Go's switch is like the one in C, C++, Java, JavaScript, and PHP,
	// except that Go only runs the selected case, not all the cases that follow.
	// In effect, the break statement that is needed at the end of each case
	// in those languages is provided automatically in Go.
	// Another important difference is that Go's switch cases need not be constants,
	// and the values involved need not be integers.
	switch os := runtime.GOOS; os { // GOOS is the running program’s operating system
	case "darwin", "ios": // means if os is either "darwin" OR "ios"
		fmt.Println("macOS Apple World") // Go has auto break!
		// fallthrough // Go down intentionally
	case "linux":
		fmt.Println("linux")
	case "windows":
		fmt.Println("windows")
	default:
		// freebsd, openbsd
		fmt.Printf("%s, \n", os)
	}

	// Alternative usage
	t := time.Now()
	fmt.Println("t: ", t, " Hour: ", t.Hour())
	// Switch without a condition is the same as switch true.
	switch { // means true
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// Define multiple variables
	switch os, user := runtime.GOOS, "genesis"; {
	case os == "linux" && user == "admin":
		fmt.Println("Linux admin")
	case os == "darwin" || os == "ios":
		fmt.Println("Apple device")
	}

	fmt.Println()
	fmt.Println()

	today := time.Now().Weekday()
	fmt.Println("When is satuday?", today, "\n\n")
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// implement a square root function:
// given a number x,
// we want to find the number z for which z² is most nearly x.
// X = 25
// Z would be 5, Z² = 5² = 25
func Sqrt(x float64) (z float64) {
	fmt.Println("x: ", x)
	fmt.Println("z: ", z)
	for n := 1; n < int(x+1); n += 1 {
		fmt.Println("n: ", n)
		if nn := float64(n * n); nn >= x {
			z = float64(n)
			break
		}
	}

	return
}

func SqrtExercise() {
	fmt.Println("Sqrt of 64: ", Sqrt(64))
	fmt.Println("Sqrt of 35: ", Sqrt(35))
	fmt.Println("Sqrt of 122: ", Sqrt(122))
	fmt.Println()
}
