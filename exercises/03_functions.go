package main

import "fmt"

// === BASIC FUNCTION ===
// func functionName(parameter type) returnType { }
func greet(name string) string {
	return "Hello, " + name + "!"
}

// === MULTIPLE RETURN VALUES (Go specialty!) ===
// In C#, you'd use out parameters or Tuple<>. Go makes it native.
// Returns (sum, product, error)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// === NAMED RETURN VALUES ===
// You can name return values. They get zero-initialized.
// Useful for clarity, but don't overuse.
func swap(x, y string) (first string, second string) {
	first = y
	second = x
	return // "naked return" — returns named values
}

// === VARIADIC FUNCTIONS ===
// Accept variable number of arguments (like params in C#)
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// === FUNCTIONS AS VALUES ===
// Go treats functions as first-class values. You can pass them around.
func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func add(x, y int) int { return x + y }
func multiply(x, y int) int { return x * y }

func main() {
	// === BASIC CALL ===
	msg := greet("Genny")
	fmt.Println(msg)

	// === MULTIPLE RETURNS ===
	// Go idiomatic: check the error second value
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Will print this
	}

	// === NAMED RETURNS ===
	first, second := swap("Go", "Rust")
	fmt.Printf("After swap: %s, %s\n", first, second)

	// === VARIADIC FUNCTIONS ===
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", total)

	// You can also pass a slice with ...
	numbers := []int{10, 20, 30}
	total = sum(numbers...) // ... unpacks the slice
	fmt.Println("Sum of slice:", total)

	// === FUNCTIONS AS VALUES ===
	result1 := applyOperation(5, 3, add)
	result2 := applyOperation(5, 3, multiply)
	fmt.Printf("5 + 3 = %d, 5 * 3 = %d\n", result1, result2)

	// Anonymous function (closure)
	subtract := func(x, y int) int {
		return x - y
	}
	result3 := applyOperation(10, 3, subtract)
	fmt.Printf("10 - 3 = %d\n", result3)

	// === EXERCISE ===
	// 1. Write a function parseAge(ageStr string) (int, error) that:
	//    - Takes a string like "25"
	//    - Returns the age as int and nil error if valid
	//    - Returns 0 and an error if it can't parse
	//    Hint: use strconv.ParseInt or strconv.Atoi
	//
	// 2. Test it with valid and invalid inputs
	//
	// 3. Write a function that takes 2 numbers and a function,
	//    and returns the result (like applyOperation above)
}

// To run:
// go run 03_functions.go
