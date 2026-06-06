# Go Quick Reference (for C# Developers)

## Essential Syntax

### Variables & Types
```go
var x int = 10           // explicit type
var y = 20               // inferred type
z := 30                  // short declaration (functions only)
const MAX = 100          // constants

// Types: int, float64, string, bool, error
// No implicit conversions! Convert explicitly: float64(x)
```

### Functions
```go
func add(a, b int) int {
    return a + b
}

// Multiple returns (Go specialty!)
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divide by zero")
    }
    return a / b, nil
}

// Variadic
func sum(nums ...int) int { /* ... */ }

// Named returns (use sparingly)
func swap(x, y string) (first, second string) {
    first, second = y, x
    return
}
```

### Control Flow
```go
if x > 0 {
    // ...
} else if x < 0 {
    // ...
} else {
    // ...
}

for i := 0; i < 10; i++ { /* */ }    // C-style
for item := range slice { /* */ }     // iteration
for {  }                               // infinite loop
switch x {
case 1: /* */
case 2, 3: /* */
default: /* */
}
```

### Structs & Methods
```go
type Person struct {
    Name string  // Exported (public)
    age  int     // Unexported (private)
}

// Value receiver (operates on copy)
func (p Person) String() string {
    return p.Name
}

// Pointer receiver (can modify)
func (p *Person) SetAge(age int) {
    p.age = age
}

// Creating structs
p := Person{Name: "Alice", age: 30}
p := Person{"Alice", 30}  // Positional (not idiomatic)
ptr := &p                  // Pointer to struct
```

### Interfaces
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Embedded interface
type ReadWriter interface {
    Reader
    Writer
}

// Any type with matching methods satisfies the interface (implicit!)
type MyReader struct{}
func (m MyReader) Read(p []byte) (n int, err error) { /* ... */ }
// MyReader now implements Reader
```

### Error Handling
```go
result, err := someFunction()
if err != nil {
    // Handle error
    return err  // or log, or retry, or use default
}

// Creating errors
return errors.New("something went wrong")
return fmt.Errorf("failed to parse: %w", underlyingErr)

// Custom errors
type CustomError struct {
    Code int
}
func (e *CustomError) Error() string {
    return fmt.Sprintf("error %d", e.Code)
}
```

### Concurrency

#### Goroutines
```go
go functionName()  // Launch concurrent function

// Wait for goroutines
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()
```

#### Channels
```go
c := make(chan int)          // Unbuffered
c := make(chan int, 5)       // Buffered (capacity 5)

c <- value                    // Send
value := <-c                  // Receive
value, ok := <-c             // Receive with ok check (useful after close)

close(c)                      // Signal no more values
for value := range c { }      // Iterate until closed

// Directional channels
func send(c chan<- int) { c <- 42 }     // Send only
func recv(c <-chan int) { <-c }         // Receive only
```

#### Select
```go
select {
case msg := <-c1:
    fmt.Println("c1:", msg)
case msg := <-c2:
    fmt.Println("c2:", msg)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
}
```

## Common Patterns

### Error Check Pattern
```go
if err != nil {
    log.Fatal(err)  // or return, or retry
}
```

### Worker Pool
```go
jobs := make(chan Job, 10)
results := make(chan Result, 10)

for i := 0; i < numWorkers; i++ {
    go worker(jobs, results)
}

// Send jobs, collect results
```

### Timeout Pattern
```go
select {
case result := <-ch:
    return result
case <-time.After(timeout):
    return nil, errors.New("timeout")
}
```

## Key Differences from C#

| C# | Go |
|---|---|
| Classes | Structs + Interfaces + Methods |
| Inheritance | Composition (embedded structs) |
| Exceptions | Error returns (no try/catch) |
| async/await | Goroutines + Channels |
| LINQ | range loops + higher-order functions |
| `null` | Zero values + explicit nil |
| `new` | `make()` for slices/maps/channels, struct literals for others |
| Generics | Limited generics (1.18+), usually use `interface{}` or code generation |

## Essential Practices

1. **Check errors immediately**: Every function that returns error, check it next line
2. **Use interfaces for abstraction**: Small interfaces (1-3 methods) are idiomatic
3. **Goroutines are cheap**: Spawn thousands if needed
4. **Channels are communication**: Prefer channels over mutexes for goroutine coordination
5. **Composition over inheritance**: Embed structs instead of inheritance hierarchies
6. **Exported vs. Unexported**: Capitalized = public, lowercase = private (package-level)
7. **Always close channels you create**: Signals to receivers that no more data is coming

## Tools

```bash
go run file.go              # Run immediately
go build                    # Build executable
go test ./...               # Run tests
go test -race ./...         # Detect race conditions
go fmt ./...                # Format code
go vet ./...                # Static analysis
go mod init modulename      # Initialize module
```

## Next: Real Projects

Once you're comfortable with these exercises:

1. **CLI Tool** — Process command-line arguments, file I/O, error handling
2. **HTTP Server** — Standard library `net/http` package
3. **Database** — `database/sql` with a driver (sqlite, postgres, mysql)
4. **Concurrency** — Worker pools, pipeline patterns
5. **Testing** — `*_test.go` files, `testing` package

All exercises are in `./exercises/`. Run with `go run XX_name.go`.
