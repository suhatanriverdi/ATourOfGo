package basics

import (
	"fmt"
	"math"
	"math/big"
	"math/cmplx"

	"reflect"

	"rsc.io/quote"
)

// Go's basic types are

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128

// Zero values
// Variables declared without an explicit initial value are given their zero value.

// The zero value is:
// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

func Basics() {
	fmt.Println("THIS IS BASICS PACKAGE/FOLDER")
	fmt.Println("Genesis Corp BASICS")
	fmt.Println(quote.Go())
	fmt.Println()

	// %T -> shows the type like int, string, bool
	// %v -> use for any value to print its default value
	fmt.Printf("Type %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value: %v\n", z, z)

	// Alternative way is to use "reflect" package
	fmt.Println("Type:", reflect.TypeOf(ToBe), "Value:", ToBe)
	fmt.Println("Type:", reflect.TypeOf(MaxInt), "Value:", MaxInt)
	fmt.Println("Type:", reflect.TypeOf(z), "Value:", z)

	// // Go 1.22+ new approach
	fmt.Println("Type:", reflect.TypeFor[bool](), "Value:", ToBe)

	// %#v
	fmt.Printf("%#v", ToBe)

	fmt.Println()
	var genesis string = "Corp"
	fmt.Println(genesis)
	fmt.Println()

	var i int
	var f float64
	var bb bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, bb, s)

	// Type Conversion
	fmt.Println()
	fmt.Println()
	var x, y int = 3, 4
	var ff float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println("ff:", ff)

	// in Go assignment between items of different type
	// requires an explicit conversion.
	var z uint = uint(ff)
	fmt.Println(x, y, z)

	// Variable's type is inferred from the value on the right hand side.
	v := 42
	var vv = 123
	fmt.Println(v, vv)

	var num int
	j := num // j is an int
	fmt.Printf("Type j is %T", j)

	fmt.Println()
	fmt.Println()
	const World = "Genesis World!"
	fmt.Println("World: ", World, " Pi: ", Pi)
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	// Constants cannot be declared using the := syntax.
	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println()
	fmt.Println()

	const (
		// Create a huge number by shifting a 1 bit left 100 places.
		// In other words, the binary number that is 1 followed by 100 zeroes.
		Big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2.
		Small = Big >> 99
	)

	fmt.Println(needInt(Small))
	// Overflow
	// fmt.Println(needInt(Big))
	bigInt := new(big.Int).Lsh(big.NewInt(1), 100)
	fmt.Println(bigInt.String())
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// In this example, untypedConst is flexible
	// while typedConst is restricted to its declared type:
	const untypedConst = 42   // Untyped integer constant
	const typedConst int = 42 // Typed integer constant

	var aa int32 = untypedConst  // Works: untypedConst adapts to int32
	var b float64 = untypedConst // Works: untypedConst adapts to float64

	// var cc int32 = typedConst    // Error: cannot use typedConst (type int) as type int32 in assignment
	// var ccf float32 = typedConst // Error: cannot use typedConst (type int) as type int32 in assignment
	var dd int = typedConst // Works: both are type int

	fmt.Println(aa, b, dd)
	fmt.Println()
}

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}
