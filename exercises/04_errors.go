package main

import (
	"errors"
	"fmt"
	"strconv"
)

// === ERROR HANDLING IN GO ===
// There are NO exceptions. Errors are VALUES returned by functions.
// This feels weird coming from C#, but it forces you to handle errors explicitly.
//
// The error type is an interface:
//   type error interface {
//       Error() string
//   }
//
// Any type with an Error() method is an error.

// === SIMPLE ERROR ===
func openFile(filename string) ([]byte, error) {
	if filename == "" {
		return nil, errors.New("filename cannot be empty")
	}
	// In real code, you'd read from disk. For now, stub it.
	return []byte("file contents"), nil
}

// === CUSTOM ERROR TYPE ===
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error in %s: %s", e.Field, e.Message)
}

func validateEmail(email string) error {
	if email == "" {
		return &ValidationError{"email", "cannot be empty"}
	}
	if !contains(email, "@") {
		return &ValidationError{"email", "must contain @"}
	}
	return nil
}

func contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// === WRAPPING ERRORS (Go 1.13+) ===
// You can wrap errors to add context while preserving the original error.
func parseUserInput(input string) (int, error) {
	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("failed to parse user input %q: %w", input, err)
	}
	return n, nil
}

// === THE GOLDEN RULE ===
// Every function that can fail returns (result, error) as its last value.
// Check it immediately.

func riskyOperation() (string, error) {
	// Simulated failure
	return "", fmt.Errorf("something went wrong")
}

func safeWrapper() string {
	result, err := riskyOperation()
	if err != nil {
		// Handle the error. Options:
		// 1. Return the error to the caller (propagate)
		// 2. Log and return a default
		// 3. Panic (rarely—only for unrecoverable errors)
		// 4. Retry

		// For now, log and return default:
		fmt.Println("Warning:", err)
		return "default value"
	}
	return result
}

func main() {
	fmt.Println("=== BASIC ERROR HANDLING ===")

	// === CHECK ERROR IMMEDIATELY ===
	data, err := openFile("myfile.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Opened file:", string(data))
	}

	// === EMPTY FILENAME ===
	data, err = openFile("")
	if err != nil {
		fmt.Println("Error:", err) // Prints: "filename cannot be empty"
	}

	fmt.Println("\n=== CUSTOM ERROR TYPE ===")

	// === VALIDATE EMAIL ===
	err = validateEmail("")
	if err != nil {
		// Type assertion to get the custom error
		if valErr, ok := err.(*ValidationError); ok {
			fmt.Printf("Field: %s, Message: %s\n", valErr.Field, valErr.Message)
		}
	}

	err = validateEmail("invalid-email")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = validateEmail("user@example.com")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Email is valid")
	}

	fmt.Println("\n=== WRAPPED ERRORS ===")

	// Wrapped errors preserve the original for inspection
	n, err := parseUserInput("not a number")
	if err != nil {
		fmt.Println("Error:", err) // Shows the full context chain
	}

	n, err = parseUserInput("42")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Parsed: %d\n", n)
	}

	fmt.Println("\n=== PROPAGATING ERRORS ===")

	result := safeWrapper()
	fmt.Println("Result:", result)

	// === EXERCISE ===
	// 1. Write a function divideNumbers(a, b string) (float64, error) that:
	//    - Parses both strings to floats
	//    - Divides a by b
	//    - Returns error if b is "0" or parsing fails
	//    Hint: use strconv.ParseFloat
	//
	// 2. Write a function that calls divideNumbers and handles both parsing errors
	//    and division-by-zero errors differently (log them differently)
	//
	// 3. Test with valid inputs, non-numeric inputs, and division by zero
}

// To run:
// go run 04_errors.go
