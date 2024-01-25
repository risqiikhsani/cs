// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// )

// type ChopS struct{ sync.Mutex }

// type PhiloChannel struct {
// 	channel   chan bool
// 	timeEaten int
// }

// type Philo struct {
// 	leftCS, rightCS *ChopS
// 	id              int
// 	pc              *PhiloChannel
// }

// func (p Philo) Eat() {

// 	for i := 0; i < 3; i++ {
// 		<-p.pc.channel

// 		if rand.Intn(2) == 0 {
// 			p.leftCS.Lock()
// 			p.rightCS.Lock()
// 		} else {
// 			p.rightCS.Lock()
// 			p.leftCS.Lock()
// 		}

// 		fmt.Println("starting to eat", p.id)

// 		fmt.Println("finishing eating", p.id)

// 		p.leftCS.Unlock()
// 		p.rightCS.Unlock()

// 		p.pc.channel <- true
// 	}
// }

// type Host struct {
// 	philosChan   map[int]*PhiloChannel
// 	philosEating [2]*PhiloChannel
// }

// func (h *Host) changeEater(i int) {
// 	minTimeEaten := 3
// 	minEatenChan := h.philosEating[i]
// 	for _, v := range h.philosChan {
// 		if v.channel != h.philosEating[0].channel && v.channel != h.philosEating[1].channel && v.timeEaten < minTimeEaten {
// 			minEatenChan = v
// 		}
// 	}
// 	h.philosEating[i] = minEatenChan
// }

// func (h *Host) StartHost(wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	h.philosEating[0] = h.philosChan[0]
// 	h.philosEating[1] = h.philosChan[2]

// 	for _, pc := range h.philosEating {
// 		pc.channel <- true
// 	}

// 	for {
// 		select {
// 		case <-h.philosEating[0].channel:
// 			h.philosEating[0].timeEaten++
// 			h.changeEater(0)
// 			if h.philosEating[0].timeEaten < 3 {
// 				h.philosEating[0].channel <- true
// 			}
// 		case <-h.philosEating[1].channel:
// 			h.philosEating[1].timeEaten++
// 			h.changeEater(1)
// 			if h.philosEating[1].timeEaten < 3 {
// 				h.philosEating[1].channel <- true
// 			}
// 		default:
// 		}

// 		allEaten := true
// 		for _, v := range h.philosChan {
// 			if v.timeEaten < 3 {
// 				allEaten = false
// 				break
// 			}
// 		}
// 		if allEaten {
// 			break
// 		}
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	host := Host{make(map[int]*PhiloChannel), [2]*PhiloChannel{}}
// 	Csticks := make([]*ChopS, 5)
// 	for i := 0; i < 5; i++ {
// 		Csticks[i] = new(ChopS)
// 	}
// 	philos := make([]*Philo, 5)
// 	for i := 0; i < 5; i++ {
// 		philoChan := &PhiloChannel{make(chan bool), 0}
// 		philos[i] = &Philo{Csticks[i], Csticks[(i+1)%5], i + 1, philoChan}
// 		host.philosChan[i] = philoChan
// 	}
// 	wg.Add(1)
// 	go host.StartHost(&wg)
// 	for i := 0; i < 5; i++ {
// 		go philos[i].Eat()
// 	}
// 	wg.Wait()

// 	for _, v := range host.philosChan {
// 		close(v.channel)
// 	}
// }
