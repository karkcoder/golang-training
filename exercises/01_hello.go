package main

import "fmt"

// This is your first Go program.
// Go programs are organized into packages. main is special—it's your entry point.
// The func main() is called when you run the program.

func main() {
	fmt.Println("Hello, Go!")

	// Try adding more lines below. For example:
	// fmt.Println("I'm learning Go")
	// fmt.Println("This is easier than I thought")
}

// To run this:
// go run 01_hello.go
//
// To build an executable:
// go build -o hello 01_hello.go
// ./hello
