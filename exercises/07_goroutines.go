package main

import (
	"fmt"
	"time"
)

// === GOROUTINES ===
// Goroutines are lightweight threads managed by the Go runtime.
// Unlike OS threads (C#'s Thread), goroutines are cheap. You can have 100,000+ easily.
//
// To launch a goroutine, just prefix a function call with "go".
// The main function waits for them... well, actually it doesn't! You have to manage that.

func sayHello(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Hello %s (%d)\n", name, i)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

func worker(id int) {
	for job := 1; job <= 5; job++ {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(200 * time.Millisecond)
	}
}

// === RACE CONDITIONS ===
// Multiple goroutines accessing shared data = race condition.
// Go has a race detector: go run -race script.go

var counter = 0 // UNSAFE to modify from multiple goroutines!

func unsafeIncrement() {
	for i := 0; i < 10000; i++ {
		counter++ // NOT safe! Try: go run -race 07_goroutines.go
	}
}

// === SYNC ===
// We'll cover channels next (preferred), but Go also has sync.Mutex

import "sync"

var (
	safeCounter = 0
	mu          sync.Mutex // Protects safeCounter
)

func safeIncrement() {
	for i := 0; i < 10000; i++ {
		mu.Lock()
		safeCounter++
		mu.Unlock()
	}
}

// === WAITING FOR GOROUTINES ===
// WaitGroup lets you wait for multiple goroutines to finish

func greetConcurrently() {
	var wg sync.WaitGroup

	names := []string{"Alice", "Bob", "Charlie"}
	for _, name := range names {
		wg.Add(1) // Increment counter
		go func(n string) {
			defer wg.Done() // Decrement counter when done
			sayHello(n)
		}(name)
	}

	wg.Wait() // Block until counter reaches 0
	fmt.Println("All greetings done!")
}

// === COORDINATING GOROUTINES WITH CHANNELS (next lesson!) ===
// Channels are how goroutines communicate. Much better than shared memory.

func main() {
	fmt.Println("=== BASIC GOROUTINE ===")

	// This launches a goroutine but main exits immediately!
	// go sayHello("World")
	//
	// Try uncommenting and running—you'll see nothing printed.
	// The program exits before the goroutine runs.

	fmt.Println("\n=== USING WaitGroup TO WAIT ===")
	greetConcurrently()

	fmt.Println("\n=== MULTIPLE WORKERS ===")

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID)
		}(i)
	}
	wg.Wait()
	fmt.Println("All workers done!")

	fmt.Println("\n=== RACE CONDITIONS ===")
	fmt.Println("Run this with: go run -race 07_goroutines.go")
	fmt.Println("(It's intentionally broken to show race detection)")

	// Uncomment to see race condition:
	// var wg2 sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	wg2.Add(1)
	// 	go func() {
	// 		defer wg2.Done()
	// 		unsafeIncrement()
	// 	}()
	// }
	// wg2.Wait()
	// fmt.Printf("Unsafe counter (should be 100000): %d\n", counter)

	fmt.Println("\n=== SAFE INCREMENT WITH MUTEX ===")
	var wg3 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			safeIncrement()
		}()
	}
	wg3.Wait()
	fmt.Printf("Safe counter (should be 100000): %d\n", safeCounter)

	fmt.Println("\n=== EXERCISE ===")
	// 1. Write a function downloadFile(url string) that sleeps for 1 second
	//    and prints "Downloaded: <url>"
	//
	// 2. Create a list of URLs and download them concurrently using goroutines
	//    and WaitGroup
	//
	// 3. Track how many downloads completed
	//
	// Bonus: Add a mutex-protected counter to track downloads
	//        (or use channels in the next lesson!)
}

// To run:
// go run 07_goroutines.go
//
// To detect race conditions:
// go run -race 07_goroutines.go
