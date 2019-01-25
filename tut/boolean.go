package main

import "fmt"

// To run this example, open the tut folder in VS Code, open a new terminal and
// type:
//
// go run boolean.go
func main() {
	// Boolean expressions can be used to evaluate common situations in
	// programming.
	//
	// Given three integers: x, y, z, we can do various comparisons between them.
	//
	// https://en.wikipedia.org/wiki/Boolean_expression
	var x int
	var y int
	var z int
	x = 1
	y = 2
	z = 2
	// fmt.Printf can print formatted data using verbs.  %d prints whole numbers,
	// %t prints 'true' or 'false' depending on the value of the boolean passed.
	fmt.Printf("Is x greater than y (%d > %d): %t\n", x, y, x > y)
	fmt.Printf("Is x less than y (%d < %d): %t\n", x, y, x < y)
	fmt.Printf("Is x less or equal to y (%d <= %d): %t\n", x, y, x <= y)
	fmt.Printf("Is z greater than or equal to y (%d >= %d): %t\n", z, y, z >= y)
	fmt.Printf("Does x have the same value as y (%d == %d): %t\n", x, y, x == y)
	fmt.Printf("Does x not have the same value as y (%d != %d): %t\n", x, y, x != y)
	fmt.Printf("Does y have the same value as z (%d == %d): %t\n", y, z, y == z)

	// We can 'chain' comparisons using OR and AND operators
	fmt.Printf("Is x greater that y OR z greater than y (%d > %d OR %d > %d): %t\n", x, y, z, y, x > y || z > y)
	fmt.Printf("Is x less than y AND z greater than y (%d < %d AND %d > %d): %t\n", x, y, z, y, x < y && z > y)
	fmt.Printf("Is x less than y AND z greater than y (%d > %d OR %d >= %d): %t\n", x, y, z, y, x > y || z >= y)

	// You can check other values too
	var a string
	var b string
	var c string
	a = "Kyle"
	b = "Kyle"
	c = "kyle"

	fmt.Printf("Does a have the same value as b (%s == %s): %t\n", a, b, a == b)
	fmt.Printf("Does a have the same value as c (%s == %s): %t\n", a, c, a == c)

	if x > 0 && c == "kyle" {
		fmt.Println(c, x)
	}
}
