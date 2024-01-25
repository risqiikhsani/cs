package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number                        int
	leftChopstick, rightChopstick *Chopstick
}

func (p Philosopher) eat(host chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ { // eat only 3 times
		<-host // acquire permission from the host

		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Printf("starting to eat %d\n", p.number)
		time.Sleep(100 * time.Millisecond) // eating time

		fmt.Printf("finishing eating %d\n", p.number)
		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()

		host <- true // signal to the host that philosopher finished eating

	}
}

func main() {
	host := make(chan bool, 2)          // host allowing 2 philosophers to eat concurrently
	chopsticks := make([]*Chopstick, 5) // make 5 chopstick
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick) // create 5 chopsticks
	}

	philosophers := make([]*Philosopher, 5) // make 5 philosophers
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%5]}
	}

	var wg sync.WaitGroup
	for _, philosopher := range philosophers {
		wg.Add(1)
		go philosopher.eat(host, &wg)
	}

	for i := 0; i < 2; i++ {
		host <- true // initially allow 2 philosophers to eat
	}

	wg.Wait()
	close(host)
}
