// package main

// import (
// 	"fmt"
// 	"sync"
// )

// const numPhilosophers = 5

// var wg sync.WaitGroup

// type Fork struct {
// 	sync.Mutex
// }

// type Philosopher struct {
// 	name                string
// 	leftFork, rightFork *Fork
// }

// func (p Philosopher) eat() {
// 	defer wg.Done()
// 	for i := 0; i < 3; i++ {
// 		p.leftFork.Lock()
// 		p.rightFork.Lock()

// 		fmt.Printf("%s is eating\n", p.name)

// 		p.leftFork.Unlock()
// 		p.rightFork.Unlock()

// 		fmt.Printf("%s is thinking\n", p.name)
// 	}
// }

// func main() {
// 	forks := make([]*Fork, numPhilosophers)
// 	for i := 0; i < numPhilosophers; i++ {
// 		forks[i] = &Fork{}
// 	}

// 	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}

// 	philosophers := make([]Philosopher, numPhilosophers)
// 	for i := 0; i < numPhilosophers; i++ {
// 		philosophers[i] = Philosopher{
// 			name:      names[i],
// 			leftFork:  forks[i],
// 			rightFork: forks[(i+1)%numPhilosophers],
// 		}
// 	}

// 	wg.Add(numPhilosophers)
// 	for i := 0; i < numPhilosophers; i++ {
// 		go philosophers[i].eat()
// 	}

// 	wg.Wait()
// }
