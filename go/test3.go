// package main

// import (
// 	"fmt"
// )

// type Animal struct {
// 	food       string
// 	locomotion string
// 	noise      string
// }

// func (a Animal) Eat() {
// 	fmt.Println(a.food)
// }

// func (a Animal) Move() {
// 	fmt.Println(a.locomotion)
// }

// func (a Animal) Speak() {
// 	fmt.Println(a.noise)
// }

// func main() {
// 	animals := map[string]Animal{
// 		"cow":   Animal{"grass", "walk", "moo"},
// 		"bird":  Animal{"worms", "fly", "peep"},
// 		"snake": Animal{"mice", "slither", "hsss"},
// 	}

// 	for {
// 		fmt.Print(">")
// 		var animalName, infoRequested string
// 		fmt.Scan(&animalName, &infoRequested)

// 		animal, found := animals[animalName]
// 		if !found {
// 			fmt.Println("Wrong input animal name. Only have cow, bird, or snake.")
// 			continue
// 		}

// 		switch infoRequested {
// 		case "eat":
// 			animal.Eat()
// 		case "move":
// 			animal.Move()
// 		case "speak":
// 			animal.Speak()
// 		default:
// 			fmt.Println("Wrong input information request. Only have eat, move, or speak.")
// 		}
// 	}
// }
