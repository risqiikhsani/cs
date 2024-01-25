// package main

// import (
// 	"fmt"
// 	"time"
// )

// // A race condition occurs when two or more goroutines access a shared resource concurrently,
// // and the final outcome depends on the timing or interleaving of their execution.
// // The result of the operations becomes unpredictable and incorrect due to unexpected interleaving of instructions.

// var counter int // Shared resource

// func incrementCounter() {
// 	for i := 0; i < 5; i++ {
// 		value := counter
// 		time.Sleep(time.Duration(value)) // Simulating some processing time
// 		value++
// 		counter = value
// 	}
// }

// func main() {
// 	fmt.Println("Starting Counter Value:", counter)
// 	go incrementCounter() // this will cause the shared resource counter value = 5
// 	go incrementCounter() // this will cause the shared resource counter value = 10

// 	time.Sleep(3 * time.Second) // Let the goroutines run for a while

// 	fmt.Println("Final Counter Value:", counter)
// }
