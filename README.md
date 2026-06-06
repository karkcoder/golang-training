# Go Training for C# Developers

A comprehensive, hands-on learning program for seasoned C# developers transitioning to Go. This course assumes you understand object-oriented programming, async patterns, and testing—and teaches you how those concepts map to Go's unique approach.

## Why Go?

Go is the language of cloud infrastructure. It powers:
- **Docker** and container orchestration
- **Kubernetes** and distributed systems
- **Cloud platforms** (AWS, GCP, Azure) have native Go support
- **Microservices** and backend APIs
- **DevOps tooling** (terraform, prometheus, consul, etc.)

As a C# developer, you already have the foundation. Go adds a different perspective on concurrency, composition, and simplicity that makes these systems easier to build.

## What's in This Repo

### 📖 Guides
- **[GET_STARTED.md](GET_STARTED.md)** — Start here. Step-by-step instructions.
- **[00_START_HERE.md](00_START_HERE.md)** — Learning path overview with module breakdown.
- **[QUICK_REFERENCE.md](QUICK_REFERENCE.md)** — Syntax cheat sheet + C# comparisons.

### 🛠️ Exercises

Eight progressive modules, each with working code and challenges:

| # | Module | Focus | Time |
|---|--------|-------|------|
| 01 | `hello.go` | Program structure, `go run` | 5 min |
| 02 | `types.go` | Variables, constants, type system | 20 min |
| 03 | `functions.go` | Functions, multiple returns, closures | 30 min |
| 04 | `errors.go` | Error handling (Go way, not exceptions) | 25 min |
| 05 | `structs.go` | Structs, methods, composition | 30 min |
| 06 | `interfaces.go` | Implicit interfaces, polymorphism | 30 min |
| 07 | `goroutines.go` | Concurrent functions, WaitGroup | 20 min |
| 08 | `channels.go` | Goroutine communication, select, patterns | 40 min |

**Total:** ~3-4 hours hands-on learning

## Quick Start

```bash
# Verify Go is installed
go version

# Run your first exercise
cd exercises
go run 02_types.go

# Work through each file in order, completing the exercises
```

## Key Concepts: C# → Go

| C# Concept | Go Equivalent | Key Difference |
|-----------|---------------|-----------------|
| Classes | Structs + Methods | No inheritance; use composition (embedding) |
| Properties | Exported fields | Capitalized = public, lowercase = private |
| try/catch/finally | if err != nil | Errors are return values, checked explicitly |
| Task/async/await | Goroutines + Channels | Lighter weight, different mental model |
| LINQ | range loops | Simpler, idiomatic Go is less fancy |
| Interface implementation | Implicit interfaces | No "implements" keyword; structural typing |
| Generics | (Limited) | Go 1.18+ has generics, but most Go code uses interface{} |
| new/constructor | Struct literals or factory functions | More explicit, less boilerplate |

## Learning Path

### Phase 1: Basics (Exercises 1-5)
Learn Go syntax, types, functions, and how to structure data with structs. This is where you'll notice Go is smaller than C#—it's by design.

**Key mindset shift:** Error handling is not exceptional. You return errors as values and check them immediately.

### Phase 2: Design Patterns (Exercise 6)
Interfaces and composition. Go doesn't have inheritance; instead, small interfaces + embedding give you flexibility without complexity.

**Key mindset shift:** No "implements" keyword. If your type has the methods, it satisfies the interface. This is called structural typing.

### Phase 3: Concurrency (Exercises 7-8)
Goroutines and channels. Go's secret sauce. Lightweight concurrency without callbacks or complex async machinery.

**Key mindset shift:** Goroutines are cheap (spawn 100,000 if needed). Channels are the preferred way to communicate between them—not shared memory with locks.

## Running Exercises

### Basic Run
```bash
go run 02_types.go
```

### Detect Race Conditions
```bash
go run -race 07_goroutines.go
go run -race 08_channels.go
```

### Format Code
```bash
go fmt
```

### Check for Issues
```bash
go vet ./...
```

## Common Patterns You'll See

### Error Handling (Go Way)
```go
result, err := someFunction()
if err != nil {
    return err  // Propagate
    // OR: log.Fatal(err)
    // OR: return defaultValue
}
```

### Goroutines + WaitGroup
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()
```

### Channels for Communication
```go
results := make(chan Result, 10)
go func() {
    results <- computeResult()
}()
result := <-results
```

### Worker Pool
```go
for i := 0; i < numWorkers; i++ {
    go worker(jobs, results)
}
for job := range jobs {
    // processed
}
```

## Tips for Success

1. **Type along.** Don't just read—run and modify the exercises.
2. **Read error messages carefully.** Go's errors are usually very helpful.
3. **Break things on purpose.** Comment out lines, change types, see what fails.
4. **Use `fmt.Printf("%#v", var)`** for debugging—shows struct internals.
5. **Trust the type system.** Go catches many bugs at compile time.
6. **Follow the error path.** Every error is a hint about what the code expects.

## After These Exercises

Once comfortable, build real projects:

- **CLI Tool** — Command-line arguments, file I/O
- **Web Server** — `net/http` package, routing
- **Database App** — `database/sql` with a driver (sqlite, postgres)
- **Concurrent System** — Worker pools, pipelines, job queues

Check out [Effective Go](https://go.dev/doc/effective_go) for idioms and [Go by Example](https://gobyexample.com/) for more patterns.

## Troubleshooting

### "command not found: go"
Install Go from https://go.dev/dl/

### "undefined: ..." error
Most likely missing an import statement. Check the exercise example.

### "race condition detected"
Multiple goroutines accessing the same variable unsafely. Use channels instead of shared memory, or sync.Mutex.

### Exercise feels unclear
- Read the comments—they explain the WHY
- Look at the examples before the exercise section
- Modify the examples and re-run to experiment

## Resources

- **Official** — https://go.dev (spec, docs, tour)
- **Effective Go** — https://go.dev/doc/effective_go
- **Go by Example** — https://gobyexample.com/
- **Standard Library** — https://pkg.go.dev/std

## Exercises At a Glance

```
exercises/
├── 01_hello.go           Basic program structure
├── 02_types.go           Variables, types, constants
├── 03_functions.go       Functions, returns, closures
├── 04_errors.go          Error handling (the Go way)
├── 05_structs.go         Structs, methods, composition
├── 06_interfaces.go      Interfaces, polymorphism
├── 07_goroutines.go      Concurrent functions, WaitGroup
└── 08_channels.go        Channels, select, worker patterns
```

Each file is standalone and runnable: `go run XX_name.go`

---

**Ready to learn Go?** Start with [GET_STARTED.md](GET_STARTED.md).

Good luck! 🚀
