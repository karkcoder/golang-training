package main

import "fmt"

// === INTERFACES IN GO ===
// Interfaces are contracts: "anything that has these methods satisfies this interface"
// You don't explicitly say "implements X". It's implicit. This is called "structural typing".
// This is VERY different from C# where you explicitly inherit/implement.
//
// Key insight: small interfaces, high reusability.

// === INTERFACE DEFINITION ===
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

// Interfaces can embed other interfaces
type ReadWriter interface {
	Reader
	Writer
}

// === IMPLEMENTING AN INTERFACE (implicitly!) ===

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// === EMPTY INTERFACE ===
// interface{} (or in Go 1.18+, any) matches everything.
// Like object in C# or interface{} in Java.

func printAnything(x interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", x, x)
}

// === TYPE ASSERTION ===
// To get the concrete value back from an interface

func describeValue(x interface{}) {
	// Type switch (like pattern matching)
	switch v := x.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case float64:
		fmt.Printf("Float: %f\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", x)
	}
}

// === POLYMORPHISM ===
func describeShape(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// === RECEIVER SEMANTICS ===
// Important: if a method has a pointer receiver (*Type), anything with that interface
// must be a pointer to pass the type check.

type Logger interface {
	Log(msg string)
}

type FileLogger struct {
	filename string
}

func (f *FileLogger) Log(msg string) {
	// Pointer receiver
	fmt.Printf("[%s] %s\n", f.filename, msg)
}

func logMessage(logger Logger, msg string) {
	logger.Log(msg)
}

func main() {
	fmt.Println("=== IMPLICIT IMPLEMENTATION ===")

	// Circle and Rectangle don't explicitly say "implements Shape"
	// But because they have Area() and Perimeter(), they satisfy Shape.
	circle := Circle{Radius: 5}
	rect := Rectangle{Width: 10, Height: 8}

	// Both can be used as Shape
	var shape Shape

	shape = circle
	fmt.Printf("Circle: ")
	describeShape(shape)

	shape = rect
	fmt.Printf("Rectangle: ")
	describeShape(shape)

	fmt.Println("\n=== POLYMORPHISM ===")

	// A slice of Shape can hold any value that satisfies Shape
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 6},
		Circle{Radius: 7},
	}

	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	fmt.Printf("Total area: %.2f\n", totalArea)

	fmt.Println("\n=== EMPTY INTERFACE ===")

	printAnything(42)
	printAnything("hello")
	printAnything(3.14)
	printAnything(circle)

	fmt.Println("\n=== TYPE ASSERTION & TYPE SWITCH ===")

	describeValue(42)
	describeValue("world")
	describeValue(2.71)
	describeValue(circle)

	// Direct type assertion (with ok check)
	val := interface{}("hello")
	if str, ok := val.(string); ok {
		fmt.Printf("It's a string: %s\n", str)
	} else {
		fmt.Println("Not a string")
	}

	fmt.Println("\n=== POINTER RECEIVERS ===")

	// FileLogger has a pointer receiver, so we need a pointer to use it as Logger
	fileLog := &FileLogger{filename: "app.log"}
	logMessage(fileLog, "Application started")

	// This would NOT work:
	// logMessage(FileLogger{filename: "app.log"}, "test") // Compiler error!

	fmt.Println("\n=== EXERCISE ===")
	// 1. Create an interface Vehicle with methods:
	//    - Speed() int
	//    - Fuel() float64
	//
	// 2. Implement Vehicle with two concrete types:
	//    - Car{speed int, fuelLevel float64}
	//    - Bike{speed int, fuelLevel float64}
	//
	// 3. Write a function that takes a []Vehicle and prints a summary
	//    (like totalArea above)
	//
	// 4. Create a slice with both Cars and Bikes, call your function
	//
	// 5. Bonus: Use a type switch to print different messages for Car vs Bike
}

// To run:
// go run 06_interfaces.go
