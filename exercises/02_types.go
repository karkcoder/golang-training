package main

import "fmt"

func main() {
	// === BASIC TYPES ===
	// Go has familiar types, but notice: no implicit conversions between types

	var age int = 30
	var height float64 = 5.9
	var name string = "Alice"
	var isLearning bool = true

	fmt.Printf("Age: %d, Height: %.1f, Name: %s, Learning: %v\n", age, height, name, isLearning)

	// === SHORT DECLARATION (inside functions only) ===
	// The := operator infers the type. This is Go's sweet spot for local variables.
	count := 42            // inferred as int
	message := "Hello"     // inferred as string
	pi := 3.14159          // inferred as float64
	active := true         // inferred as bool

	fmt.Println(count, message, pi, active)

	// === ZERO VALUES ===
	// Unlike C#, Go initializes vars to zero values if you don't assign:
	// int=0, string="", bool=false, pointers=nil
	var uninitialized int
	var emptyString string
	var notSet bool

	fmt.Printf("int zero value: %d, string zero value: %q, bool zero value: %v\n", uninitialized, emptyString, notSet)

	// === TYPE CONVERSION ===
	// Go is strict about types. You must convert explicitly.
	var x int = 10
	var y float64 = float64(x) // explicit conversion needed
	fmt.Printf("x as float64: %f\n", y)

	// === CONSTANTS ===
	const maxConnections = 100
	const greeting = "Welcome to Go"
	const pi2 = 3.14159

	fmt.Println(maxConnections, greeting, pi2)

	// Constants are computed at compile time. They must be simple expressions.
	// const futureDate = time.Now() // ERROR: not a constant

	// === EXERCISE ===
	// 1. Declare a variable for your current project name using :=
	// 2. Declare a variable for lines of code you've written (use int)
	// 3. Print them both using fmt.Printf with %s and %d
	//
	// Example:
	// projectName := "MyGoApp"
	// linesOfCode := 1500
	// fmt.Printf("Project %s has %d lines\n", projectName, linesOfCode)
}

// To run:
// go run 02_types.go
