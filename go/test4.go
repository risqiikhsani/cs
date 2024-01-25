// package main

// import (
// 	"errors"
// 	"fmt"
// )

// type Animal interface {
// 	Eat()
// 	Move()
// 	Speak()
// }

// type Cow struct{}

// func (c Cow) Eat() {
// 	fmt.Println("grass")
// }

// func (c Cow) Move() {
// 	fmt.Println("walk")
// }

// func (c Cow) Speak() {
// 	fmt.Println("moo")
// }

// type Bird struct{}

// func (b Bird) Eat() {
// 	fmt.Println("worms")
// }

// func (b Bird) Move() {
// 	fmt.Println("fly")
// }

// func (b Bird) Speak() {
// 	fmt.Println("peep")
// }

// type Snake struct{}

// func (s Snake) Eat() {
// 	fmt.Println("mice")
// }

// func (s Snake) Move() {
// 	fmt.Println("slither")
// }

// func (s Snake) Speak() {
// 	fmt.Println("hsss")
// }

// func main() {
// 	animals := make(map[string]Animal)

// 	for {
// 		var command string
// 		var string2 string
// 		var string3 string
// 		fmt.Print("> ")
// 		fmt.Scan(&command, &string2, &string3)

// 		switch command {
// 		case "newanimal":
// 			animalType := string3
// 			newAnimal, err := createAnimal(animalType)
// 			if err != nil {
// 				fmt.Println("Invalid input , string 2 (animal name) , string 3 (animal type)")
// 				fmt.Println(err)
// 				continue
// 			}
// 			animals[string2] = newAnimal
// 			fmt.Println("Created it!")

// 		case "query":
// 			requestedstring3 := string3
// 			animal, ok := animals[string2]
// 			if !ok {
// 				fmt.Println("Animal not found")
// 				continue
// 			}
// 			switch requestedstring3 {
// 			case "eat":
// 				animal.Eat()
// 			case "move":
// 				animal.Move()
// 			case "speak":
// 				animal.Speak()
// 			default:
// 				fmt.Println("Invalid string3 must be eat, move, or speak")
// 			}
// 		default:
// 			fmt.Println("Invalid command")
// 		}
// 	}
// }

// func createAnimal(animalType string) (Animal, error) {
// 	switch animalType {
// 	case "cow":
// 		return Cow{}, nil
// 	case "bird":
// 		return Bird{}, nil
// 	case "snake":
// 		return Snake{}, nil
// 	default:
// 		return nil, errors.New("Invalid animal type")
// 	}
// }
