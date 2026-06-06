package main

import (
	"fmt"
	"time"
)

// === CHANNELS ===
// Channels are the Go way to let goroutines communicate safely.
// They're like message queues. Send on one end, receive on the other.
//
// vs. C#: channels are Go's alternative to lock-based synchronization.
// Motto: "Don't communicate by sharing memory; share memory by communicating."

// === BASIC CHANNEL ===
// chan T is a channel of type T
// Send: c <- value
// Receive: value := <-c

func producer(c chan string) {
	c <- "Hello"
	c <- "from"
	c <- "goroutine"
	close(c) // Signal that no more values will be sent
}

func consumer(c chan string) {
	for msg := range c {
		fmt.Println("Received:", msg)
	}
	fmt.Println("Channel closed")
}

// === BUFFERED CHANNELS ===
// Unbuffered: send blocks until someone receives
// Buffered: send blocks only when buffer is full

func produceWithBuffer() {
	// Buffer size 2: can hold 2 values before blocking
	c := make(chan int, 2)

	c <- 1
	c <- 2
	// c <- 3 // Would block! Buffer is full.

	val1 := <-c
	val2 := <-c
	fmt.Printf("Got: %d, %d\n", val1, val2)
}

// === DIRECTIONAL CHANNELS ===
// chan<- T sends only
// <-chan T receives only
// This is a safety feature

func send(c chan<- int) {
	c <- 42
	// value := <-c // Compiler error! Can't receive on send-only channel
}

func receive(c <-chan int) {
	value := <-c
	fmt.Println("Received:", value)
	// c <- 99 // Compiler error! Can't send on receive-only channel
}

// === SELECT ===
// Wait on multiple channels, handle whichever is ready

func selectExample() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("c1:", msg1)
		case msg2 := <-c2:
			fmt.Println("c2:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout")
		}
	}
}

// === TIMEOUT PATTERN ===
func fetchWithTimeout(duration time.Duration) (string, error) {
	c := make(chan string)

	go func() {
		time.Sleep(duration / 2) // Simulate work
		c <- "result"
	}()

	select {
	case result := <-c:
		return result, nil
	case <-time.After(duration):
		return "", fmt.Errorf("timeout after %v", duration)
	}
}

// === WORKER POOL (common pattern) ===
// Multiple workers listening on a job channel

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2
	}
}

func workerPool() {
	numWorkers := 3
	numJobs := 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Signal workers that no more jobs

	// Collect results
	for i := 0; i < numJobs; i++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}

func main() {
	fmt.Println("=== BASIC CHANNELS ===")

	c := make(chan string)
	go producer(c)
	consumer(c)

	fmt.Println("\n=== BUFFERED CHANNELS ===")
	produceWithBuffer()

	fmt.Println("\n=== DIRECTIONAL CHANNELS ===")
	c3 := make(chan int)
	go send(c3)
	receive(c3)

	fmt.Println("\n=== SELECT ===")
	selectExample()

	fmt.Println("\n=== TIMEOUT ===")
	result, err := fetchWithTimeout(2 * time.Second)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Got:", result)
	}

	fmt.Println("\n=== WORKER POOL ===")
	workerPool()

	fmt.Println("\n=== EXERCISE ===")
	// 1. Write a function that creates a channel and launches 5 goroutines
	//    Each goroutine sends a number (1-5) to the channel
	//
	// 2. The main function receives all values and prints them
	//
	// 3. Bonus: Use a buffered channel and experiment with different buffer sizes
	//
	// 4. Bonus 2: Implement a simple pipeline:
	//    - producer goroutine sends numbers 1-10 to a channel
	//    - processor goroutine reads, squares them, sends to output channel
	//    - main reads from output and prints
}

// To run:
// go run 08_channels.go
