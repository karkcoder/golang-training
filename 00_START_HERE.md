# Go Learning Path for C# Developers

Welcome! You're going from C# to Go. This path builds from zero to production-ready.

## Before You Start

```bash
# Verify Go is installed
go version

# Create a workspace for exercises
mkdir -p $HOME/go-learn
cd $HOME/go-learn
```

## Module 1: Basics (2-3 hours)

Start with the absolute fundamentals. Go is small and readable—you'll see it all.

### 1.1 Your First Program
[→ Exercise: hello.go](./exercises/01_hello.go)

### 1.2 Types & Variables
[→ Exercise: types.go](./exercises/02_types.go)

### 1.3 Functions & Multiple Returns
[→ Exercise: functions.go](./exercises/03_functions.go)

### 1.4 Error Handling (The Go Way)
[→ Exercise: errors.go](./exercises/04_errors.go)

### 1.5 Structs & Methods
[→ Exercise: structs.go](./exercises/05_structs.go)

---

## Module 2: Interfaces & Composition (2 hours)

This is where Go clicks. Forget inheritance.

[→ Exercise: interfaces.go](./exercises/06_interfaces.go)

---

## Module 3: Concurrency (3-4 hours)

Go's superpower. Goroutines make async easy.

[→ Exercise: goroutines.go](./exercises/07_goroutines.go)
[→ Exercise: channels.go](./exercises/08_channels.go)

---

## Module 4: Building Real Things (4+ hours)

Apply it to actual projects.

[→ Project: CLI Tool](./projects/01_weather_cli/)
[→ Project: Web Server](./projects/02_simple_web/)

---

## Keep This Nearby

- **Effective Go** (official): https://go.dev/doc/effective_go
- **Go Proverbs** (Jay Pinkerton): Philosophical guide to Go style

---

**Next Step:** Open `exercises/01_hello.go` and run it.
