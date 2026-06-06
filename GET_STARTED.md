# Get Started: Your Go Learning Journey

You're all set! Here's your next step-by-step.

## Step 1: Verify Go Installation

```bash
go version
# Should print something like: go version go1.21.0 linux/amd64
```

## Step 2: Run Your First Exercise

```bash
cd /home/firestorm/source/golang/exercises
go run 01_hello.go
```

You should see:
```
Hello, Go!
```

## Step 3: Work Through Each Exercise

Follow this order. Each builds on the previous:

1. **01_hello.go** — "Hello, Go!" (5 min)
   - Shows basic program structure
   - `go run` command
   
2. **02_types.go** — Variables and types (20 min)
   - Types, `:=` operator, constants
   - **Exercise:** Declare variables, print them
   
3. **03_functions.go** — Functions and multiple returns (30 min)
   - Function definition, multiple returns
   - Variadic functions, functions as values
   - **Exercise:** Write parseAge, division functions
   
4. **04_errors.go** — Error handling (25 min)
   - Go's error model (not exceptions!)
   - Custom error types, wrapping errors
   - **Exercise:** parseUserInput with error wrapping
   
5. **05_structs.go** — Structs and methods (30 min)
   - Defining structs, methods, receivers
   - Composition (embedded structs)
   - **Exercise:** Product struct with Discount and Sell methods
   
6. **06_interfaces.go** — Interfaces (30 min)
   - Implicit implementation (no "implements" keyword!)
   - Polymorphism, type assertions
   - **Exercise:** Vehicle interface with Car/Bike
   
7. **07_goroutines.go** — Goroutines basics (20 min)
   - Launching concurrent functions
   - sync.WaitGroup, race conditions
   - **Exercise:** Download files concurrently
   
8. **08_channels.go** — Channels (40 min)
   - Channel communication between goroutines
   - Buffered channels, select, timeouts
   - **Exercise:** Producer-consumer pipeline

**Total estimated time:** ~3-4 hours hands-on

## Step 4: Use the Quick Reference

Keep [QUICK_REFERENCE.md](QUICK_REFERENCE.md) open as you work. It has:
- Syntax quick reference
- Common patterns
- Key differences from C#

## Step 5: Experiments

Each exercise has an **Exercise** section at the bottom. Don't skip these!

- Try modifying the examples
- Break things and fix them
- Use `go run` to test immediately

## Step 6: Next Phase (after exercises)

Once you're comfortable, move to real projects:

1. **CLI Tool** — Read command-line args, process files
2. **Web Server** — `net/http` package
3. **Testing** — `*_test.go` files

We'll set those up once you finish the exercises.

## Tips

### Running with Race Detection
Go can detect race conditions (multiple goroutines accessing shared memory unsafely):

```bash
go run -race 07_goroutines.go
go run -race 08_channels.go
```

### Formatting Your Code
Go has one standard format (no style debates):

```bash
go fmt ./...  # Format all files
```

### Debugging
- Use `fmt.Printf()` for debugging (no debugger needed in Go)
- Use `fmt.Printf("%#v", variable)` to see struct internals
- Use `fmt.Printf("%T", variable)` to see a variable's type

### Getting Unstuck
- Read the error message carefully (Go's error messages are usually very helpful)
- Check the exercise comments—they often hint at solutions
- Use https://go.dev/ref/spec for language details
- Try `go doc fmt.Printf` in the terminal to see documentation

## Structure

```
/home/firestorm/source/golang/
├── 00_START_HERE.md           ← You are here
├── GET_STARTED.md             ← This file
├── QUICK_REFERENCE.md         ← Keep this handy
├── exercises/
│   ├── 01_hello.go
│   ├── 02_types.go
│   ├── 03_functions.go
│   ├── 04_errors.go
│   ├── 05_structs.go
│   ├── 06_interfaces.go
│   ├── 07_goroutines.go
│   └── 08_channels.go
└── projects/
    └── (coming next, after exercises)
```

---

**Ready?** Let's go!

```bash
cd /home/firestorm/source/golang/exercises
go run 02_types.go
```

After each exercise, let me know what you'd like to tackle next or if you hit any blockers. I can help explain concepts, debug, or provide more examples.
