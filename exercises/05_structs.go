package main

import "fmt"

// === STRUCTS (like classes without the inheritance) ===
// Struct fields are capitalized (exported) or lowercase (private)
// This is Go's access control mechanism.

type Person struct {
	Name    string
	Age     int
	Email   string
	isAdmin bool // lowercase = private to this package
}

// === METHODS ===
// Methods are functions with a receiver. The receiver appears between func and the method name.
// This is how Go does "instance methods" without classes.

// Value receiver: operates on a copy
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s", p.Name)
}

// Pointer receiver: can modify the original
func (p *Person) HaveBirthday() {
	p.Age++
}

func (p *Person) SetAdmin(admin bool) {
	p.isAdmin = admin
}

func (p Person) IsAdmin() bool {
	return p.isAdmin
}

// === EMBEDDED STRUCTS (composition over inheritance) ===
// Instead of inheriting from a base class, embed a struct.

type Address struct {
	Street string
	City   string
	ZIP    string
}

type Employee struct {
	Person  // Embedded struct (anonymous field)
	Address // Another embedded struct
	Salary  int
}

// === STRUCT TAGS ===
// Metadata attached to fields, used by other packages (JSON, database, etc.)
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"` // omitempty: skip if empty in JSON
}

func main() {
	fmt.Println("=== CREATING STRUCTS ===")

	// === BASIC CONSTRUCTION ===
	person1 := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	fmt.Println(person1)

	// Zero values for unspecified fields
	person2 := Person{Name: "Bob"}
	fmt.Printf("Bob: %s, Age: %d, isAdmin: %v\n", person2.Name, person2.Age, person2.isAdmin)

	// All-positional (not idiomatic unless struct is tiny)
	person3 := Person{"Charlie", 25, "charlie@example.com", false}
	fmt.Println(person3)

	// Pointers to structs
	personPtr := &person1
	personPtr.Age = 31 // Go auto-dereferences pointers to struct fields
	fmt.Println("After birthday:", person1)

	fmt.Println("\n=== METHODS ===")

	// Calling methods
	greeting := person1.Greet()
	fmt.Println(greeting)

	// Value receiver (operates on a copy)
	person2.HaveBirthday() // Doesn't modify person2!
	fmt.Printf("After birthday (value receiver): %v\n", person2.Age)

	// Pointer receiver (modifies original)
	personPtr.HaveBirthday()
	fmt.Printf("After birthday (pointer receiver): %v\n", person1.Age)

	personPtr.SetAdmin(true)
	fmt.Printf("Admin status: %v\n", personPtr.IsAdmin())

	fmt.Println("\n=== EMBEDDED STRUCTS (COMPOSITION) ===")

	// Employee embeds Person and Address, so it has all their fields
	emp := Employee{
		Person: Person{
			Name:  "Diana",
			Age:   28,
			Email: "diana@company.com",
		},
		Address: Address{
			Street: "123 Main St",
			City:   "San Francisco",
			ZIP:    "94105",
		},
		Salary: 120000,
	}

	// Access embedded fields directly (no need for emp.Person.Name)
	fmt.Printf("Employee: %s, %s, %s\n", emp.Name, emp.City, emp.Email)

	// Calling inherited methods
	fmt.Println(emp.Greet())

	// Can still be explicit if there's ambiguity
	fmt.Printf("Full title: %s at %s, %s\n", emp.Person.Name, emp.Address.Street, emp.Address.City)

	fmt.Println("\n=== STRUCT TAGS (JSON) ===")

	// Tags are usually used by other packages for marshaling
	user := User{ID: 1, Name: "Eve", Email: "eve@example.com"}
	fmt.Printf("%+v\n", user) // %+v includes field names

	fmt.Println("\n=== EXERCISE ===")
	// 1. Create a Product struct with:
	//    - Name (string)
	//    - Price (float64)
	//    - Stock (int)
	//
	// 2. Write a method Discount(percent float64) float64 that returns the discounted price
	//
	// 3. Write a pointer method Sell(quantity int) error that:
	//    - Decreases stock by quantity
	//    - Returns error if quantity > stock
	//    - Otherwise returns nil
	//
	// 4. Create a product, call both methods, and verify it works
}

// To run:
// go run 05_structs.go
