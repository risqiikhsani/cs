// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"runtime"
// 	"strconv"
// 	"time"
// )

// func f(n int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(n, ":", i)

// 		amt := time.Duration(rand.Intn(250))
// 		time.Sleep(time.Millisecond * amt)
// 	}
// }

// func ping(c chan<- string) { // chan<- instead of chan restrict it to send
// 	for i := 0; ; i++ {
// 		c <- "ping" + " " + strconv.Itoa(i)
// 	}
// }

// func pong(c chan<- string) { // chan<- instead of chan restrict it to send
// 	for i := 0; ; i++ {
// 		c <- "pong" + " " + strconv.Itoa(i)
// 	}
// }

// func printer(c chan string) {
// 	for {
// 		msg := <-c
// 		fmt.Println(msg)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func sum(numbers []int, ch chan int) {
// 	sum := 0
// 	for _, num := range numbers {
// 		sum += num
// 	}
// 	ch <- sum
// }

// func worker(in <-chan int, out chan<- int) {
// 	for val := range in {
// 		result := val * 2
// 		out <- result
// 	}
// }

// func sender(ch chan<- int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }

// func main() {
// 	// var c chan string = make(chan string)
// 	// go ping(c)
// 	// go pong(c)
// 	// go printer(c)

// 	// example 1
// 	//https://www.golang-book.com/books/intro/10
// 	// c1 := make(chan string)
// 	// c2 := make(chan string)

// 	// go func() {
// 	// 	for {
// 	// 		c1 <- "from 1"
// 	// 		time.Sleep(time.Second * 1)
// 	// 	}
// 	// }()

// 	// go func() {
// 	// 	for {
// 	// 		c2 <- "from 2"
// 	// 		time.Sleep(time.Second * 1)
// 	// 	}
// 	// }()

// 	// go func() {
// 	// 	for {
// 	// 		select {
// 	// 		case msg1 := <-c1:
// 	// 			fmt.Println(msg1)
// 	// 		case msg2 := <-c2:
// 	// 			fmt.Println(msg2)
// 	// 		case <-time.After(time.Second):
// 	// 			fmt.Println("timeout")
// 	// 			// default:
// 	// 			// 	fmt.Println("nothing ready")
// 	// 		}
// 	// 	}
// 	// }()

// 	// example 2
// 	// numbers := []int{1, 2, 3, 4, 5}
// 	// ch := make(chan int)
// 	// go sum(numbers[:len(numbers)/2], ch)
// 	// go sum(numbers[len(numbers)/2:], ch)
// 	// p1, p2 := <-ch, <-ch

// 	// totalSum := p1 + p2
// 	// fmt.Println(totalSum)

// 	// go f(1)
// 	// f(0)

// 	// example 3
// 	// data := []int{1, 2, 3, 4, 5, 6}

// 	// inputCh := make(chan int)
// 	// outputCh := make(chan int)

// 	// for i := 0; i < 3; i++ {
// 	// 	// create 3 worker
// 	// 	go worker(inputCh, outputCh)
// 	// }

// 	// go func() {
// 	// 	for i := 0; i < len(data); i++ {
// 	// 		fmt.Println(<-outputCh)
// 	// 	}
// 	// }()

// 	// for _, d := range data {
// 	// 	inputCh <- d
// 	// }

// 	// close(inputCh)

// 	// example 4 select
// 	// ch1 := make(chan string)
// 	// ch2 := make(chan string)

// 	// go func() {
// 	// 	time.Sleep(time.Second)
// 	// 	ch1 <- "one"
// 	// }()

// 	// go func() {
// 	// 	time.Sleep(time.Second)
// 	// 	ch2 <- "two"
// 	// }()

// 	// for i := 0; i < 2; i++ {
// 	// 	select {
// 	// 	case mg1 := <-ch1:
// 	// 		fmt.Println("Received : ", mg1)
// 	// 	case mg2 := <-ch2:
// 	// 		fmt.Println("Received : ", mg2)
// 	// 	}
// 	// }

// 	//example 5 buffered channel
// 	ch := make(chan int, 3) // Buffered channel with a capacity of 3
// 	go sender(ch)

// 	for val := range ch {
// 		fmt.Println("Received:", val)
// 	}

// 	numCPU := runtime.NumCPU()
// 	fmt.Println("Number of CPU cores:", numCPU)

// 	numThreads := runtime.GOMAXPROCS(0)
// 	fmt.Println("Number of OS threads used by Go:", numThreads)

// 	var input string
// 	fmt.Scanln(&input)
// }
