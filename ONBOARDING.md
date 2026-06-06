# Go Onboarding Guide

Welcome to Go! This guide will get you set up and ready to start learning.

## ✅ Prerequisites Check

You should have:
- A terminal/command line
- A text editor or IDE (VS Code recommended)
- ~4-5 hours for the learning exercises
- Curiosity about how Go differs from C#

## 📦 Go Installation

### Verify Installation

```bash
go version
```

Should print something like:
```
go version go1.23.0 linux/amd64
```

If it doesn't work, Go needs to be installed. See **Installation** below.

### Installation

If Go is not installed, follow these steps:

#### Linux/macOS
```bash
# Download the latest Go binary
curl -L https://go.dev/dl/go1.23.0.linux-amd64.tar.gz -o go.tar.gz

# Extract to your home directory
mkdir -p ~/.local
tar -xzf go.tar.gz -C ~/.local
rm go.tar.gz

# Add to PATH
echo 'export PATH=$HOME/.local/go/bin:$PATH' >> ~/.bashrc
source ~/.bashrc
```

#### Windows
1. Download from https://go.dev/dl/
2. Run the installer
3. Verify: `go version` in PowerShell

#### Verify Installation
```bash
go version
go env GOPATH
```

## 🗂️ Directory Structure

After cloning/setting up this repo, you'll have:

```
golang-training/
├── README.md                # Main overview
├── ONBOARDING.md           # This file
├── GET_STARTED.md          # Quick start instructions
├── QUICK_REFERENCE.md      # Syntax cheat sheet
├── exercises/
│   ├── 01_hello.go
│   ├── 02_types.go
│   ├── 03_functions.go
│   ├── 04_errors.go
│   ├── 05_structs.go
│   ├── 06_interfaces.go
│   ├── 07_goroutines.go
│   └── 08_channels.go
└── projects/               # Coming later
```

## 🎯 Your First Steps

### 1. Navigate to the Repo
```bash
cd ~/source/golang  # Or wherever you cloned it
```

### 2. Run Your First Program
```bash
cd exercises
go run 01_hello.go
```

You should see:
```
Hello, Go!
```

### 3. Read the Code
Open `01_hello.go` in your editor. Read the comments. Run it.

### 4. Progress to the Next Exercise
```bash
go run 02_types.go
```

Read, experiment, complete the exercise at the bottom.

## 🛠️ Tools & Editor Setup

### VS Code (Recommended)

Install the **Go** extension:
1. Open Extensions (Cmd+Shift+X / Ctrl+Shift+X)
2. Search for "Go"
3. Install the official Go extension by Go Team at Google

You get:
- Syntax highlighting
- Code completion
- Debugging
- Testing integration
- Linting (go vet)

### Command Line Tools

Essential commands you'll use:

```bash
go run file.go              # Run immediately
go build                    # Compile to binary
go test ./...               # Run tests
go test -race ./...         # Detect race conditions
go fmt ./...                # Format code (required!)
go vet ./...                # Static analysis
go doc fmt.Println          # See documentation
go mod init modulename      # Start a new project
go mod tidy                 # Clean up dependencies
```

### Shell Completion

For bash:
```bash
go run github.com/posener/complete/gocomplete@latest -install=true
```

For zsh:
```bash
mkdir -p ~/.zsh/completions
go run github.com/posener/complete/gocomplete@latest -zsh > ~/.zsh/completions/_go
```

## 📚 Learning Structure

### Phase 1: Basics (2 hours)
**Files: 01_hello.go → 05_structs.go**

Learn:
- Program structure
- Types and variables
- Functions and returns
- Error handling (the Go way!)
- Structs and methods

### Phase 2: Design (30 min)
**File: 06_interfaces.go**

Learn:
- Implicit interfaces (no "implements" keyword)
- Polymorphism
- Type assertions
- Composition patterns

### Phase 3: Concurrency (1.5 hours)
**Files: 07_goroutines.go → 08_channels.go**

Learn:
- Lightweight goroutines
- Communication via channels
- Worker pools
- Race detection

**Total estimated time: 3.5-4 hours**

## 🚀 Workflow

For each exercise:

```
1. Read the comments and examples in the file
2. Run it: go run XX_name.go
3. Modify the examples (break things on purpose!)
4. Complete the exercise at the bottom
5. Move to the next file
```

### Example Workflow for 02_types.go

```bash
# 1. Run the exercise
go run 02_types.go

# 2. Open in your editor and read comments
code 02_types.go  # VS Code

# 3. Modify an example—change a value, re-run
go run 02_types.go

# 4. Complete the exercise at the bottom (uncomment or add code)
go run 02_types.go

# 5. If it doesn't work, read the error carefully
# 6. Move on
```

## 🐛 Debugging Tips

### Run with Verbose Output
```bash
go run -v file.go
```

### Check for Race Conditions
```bash
go run -race file.go
```

### Debug with Print Statements
```go
fmt.Printf("%#v\n", myStruct)     // Show struct with field names
fmt.Printf("%T\n", variable)      // Show type of variable
fmt.Printf("%[1]T %[1]v\n", x)    // Type and value together
```

### Get Help on a Function
```bash
go doc fmt.Println
go doc sync.WaitGroup
```

## 📖 Reading Code

Go code is intentionally readable. Key things to look for:

1. **Function signatures** — Tell you what it takes and returns
2. **Error returns** — Check if a function returns an error
3. **Receivers** — See if a method takes a value or pointer receiver
4. **Interface satisfaction** — Look for the methods a type implements
5. **Comments** — Why something is done, not what it does

## 🔗 Key Resources

- **Official Go Tour** — https://go.tour/welcome/1
- **Effective Go** — https://go.dev/doc/effective_go (read after exercises)
- **Go by Example** — https://gobyexample.com/
- **Standard Library Docs** — https://pkg.go.dev/std
- **Go Spec** — https://go.dev/ref/spec (reference, not tutorial)

## ❓ Common Questions

### Q: Do I need to install anything else?
**A:** No, Go is self-contained. Just the one download.

### Q: Can I use an IDE other than VS Code?
**A:** Yes. GoLand, Vim, Emacs, and others have Go support. VS Code is just recommended for beginners.

### Q: How much memory does Go use?
**A:** Very little. Go is efficient.

### Q: Can I run Go on Windows/macOS/Linux?
**A:** Yes. Go runs everywhere. These exercises work on all platforms.

### Q: What if I'm stuck on an exercise?
**A:** 
1. Re-read the comments in the file
2. Check the error message (Go errors are usually clear)
3. Look at the examples before the exercise
4. Try modifying an example to understand it better
5. Skip it and come back later

### Q: How long does learning Go take?
**A:** These exercises: 4 hours. To be productive: a few weeks of building small projects. To really master Go's idioms: 6-12 months of serious use.

## 🎓 After Onboarding

After completing all 8 exercises:

1. **Review** [QUICK_REFERENCE.md](QUICK_REFERENCE.md) for syntax you want to remember
2. **Build** a small project (CLI tool, web server, file processor)
3. **Read** [Effective Go](https://go.dev/doc/effective_go) for idioms
4. **Explore** the [standard library](https://pkg.go.dev/std) — you'll be surprised how much is there

## 📝 Notes

### Go Versioning
These exercises use Go 1.18+. If you have an older version, upgrade:
```bash
go install golang.org/dl/go1.23.0@latest
~/go/bin/go1.23.0 download
```

### Module System
Go 1.11+ uses modules for dependency management. These exercises don't need external dependencies, so you won't touch `go.mod` yet.

### IDE Tips
- **Gofmt on save** — Enable automatic formatting (Go style is enforced)
- **Lint on save** — Use golangci-lint for static analysis
- **IntelliSense** — Works great once the Go extension loads

## ✨ Next Steps

1. Verify Go is installed: `go version`
2. Read [GET_STARTED.md](GET_STARTED.md)
3. Run: `go run exercises/02_types.go`
4. Start learning!

---

**Questions?** The Go community is incredibly helpful:
- **Stack Overflow** — Tag: `go` or `golang`
- **Go Forum** — https://forum.golangbridge.org
- **Gophers Slack** — https://invite.slack.golangbridge.org/

Good luck! 🎉
