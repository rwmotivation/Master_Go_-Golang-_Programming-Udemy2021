package main

import "fmt"

func main() {
	a, b := 10, 5.5

	//** ARITHMETIC OPERATORS **//
	//  +       sum
	// -        difference
	// *        product
	// /        quotient
	// %        remainder
	// there is no power operator in Go. Use math.Pow(a, b) for

	fmt.Println(a + 5)
	fmt.Println(3.1 - b)
	fmt.Println(a * a)
	fmt.Println(a / a)
	fmt.Println(11 / 5)

	//go is strong Typed Language

	fmt.Println(a * int(b))
	fmt.Println(float64(a) * b)

	//incDEC Statements
	//THE  "++" statements increment or decrement their operands by the untyped constant 1.
	x := 10
	x++ // x is 11 same as x += 1
	x-- //x is 10  same as x -= 1

	//** ASSIGNMENT OPERATORS **//
	//  =   (simple assignment)
	// +=   (increment assignment)
	// -=   (decrement assignment)
	// *=   (multiplication assignment)
	// /=   (division assignment)
	// %=   (modulus assignment)

	a = 10
	a += 2 // => a is 12
	a -= 3 // => a is 9
	a *= 2 // => a is 18
	a /= 3 // => a is 6
	a %= 5 // => a is 1

	fmt.Println(5 == 6)   // => false
	fmt.Println(5 != 6)   // => true
	fmt.Println(10 > 10)  // => false
	fmt.Println(10 >= 10) // => true
	fmt.Println(5 < 5)    // => false
	fmt.Println(5 <= 5)   // => true

	//** LOGICAL OPERATORS **//
	// &&       logical and
	// ||       logical or
	// !        logical negation

	fmt.Println(0 < 2 && 4 > 1) // => true
	fmt.Println(1 > 5 || 4 > 5) // => false
	fmt.Println(!(1 > 2))       // => true

}
