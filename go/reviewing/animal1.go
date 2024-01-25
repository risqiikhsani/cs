// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// type animal struct {
// 	food       string
// 	locomotion string
// 	noise      string
// }

// type Cow struct {
// 	food       string
// 	locomotion string
// 	noise      string
// }

// type Bird struct {
// 	food       string
// 	locomotion string
// 	noise      string
// }

// type Snake struct {
// 	food       string
// 	locomotion string
// 	noise      string
// }

// type Animal interface {
// 	Eat() string
// 	Move() string
// 	Speak() string
// }

// func (c Cow) Eat() string {
// 	return c.food
// }

// func (c Cow) Move() string {
// 	return c.locomotion
// }

// func (c Cow) Speak() string {
// 	return c.noise
// }

// func (b Bird) Eat() string {
// 	return b.food
// }

// func (b Bird) Move() string {
// 	return b.locomotion
// }

// func (b Bird) Speak() string {
// 	return b.noise
// }

// func (s Snake) Eat() string {
// 	return s.food
// }

// func (s Snake) Move() string {
// 	return s.locomotion
// }

// func (s Snake) Speak() string {
// 	return s.noise
// }

// var animalData = map[string]Animal{
// 	"cow": Cow{
// 		food:       "grass",
// 		locomotion: "walk",
// 		noise:      "moo",
// 	},
// 	"bird": Bird{
// 		food:       "worms",
// 		locomotion: "fly",
// 		noise:      "peep",
// 	},
// 	"snake": Snake{
// 		food:       "mice",
// 		locomotion: "slither",
// 		noise:      "hsss",
// 	},
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	if err := scanner.Err(); err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Print(">")
// 	for scanner.Scan() {
// 		input := scanner.Text()
// 		data := strings.Split(input, " ")
// 		data[0] = strings.ToLower(data[0])
// 		data[1] = strings.ToLower(data[1])
// 		data[2] = strings.ToLower(data[2])

// 		switch data[0] {
// 		case "newanimal":
// 			switch data[2] {
// 			case "cow":
// 				animalData[data[1]] = Cow{
// 					food:       "grass",
// 					locomotion: "walk",
// 					noise:      "moo",
// 				}
// 				fmt.Println("Created it!")
// 			case "bird":
// 				animalData[data[1]] = Bird{
// 					food:       "worms",
// 					locomotion: "fly",
// 					noise:      "peep",
// 				}
// 				fmt.Println("Created it!")
// 			case "snake":
// 				animalData[data[1]] = Snake{
// 					food:       "mice",
// 					locomotion: "slither",
// 					noise:      "hsss",
// 				}
// 				fmt.Println("Created it!")
// 			default:
// 				fmt.Println("Invalid input")
// 			}
// 		case "query":
// 			switch data[2] {
// 			case "eat":
// 				fmt.Printf("%s \n", animalData[data[1]].Eat())
// 			case "move":
// 				fmt.Printf("%s \n", animalData[data[1]].Move())
// 			case "speak":
// 				fmt.Printf("%s \n", animalData[data[1]].Speak())
// 			default:
// 				fmt.Println("Invalid input")
// 			}

// 		default:
// 			fmt.Println("Invalid input")
// 		}
// 		fmt.Print(">")
// 	}
// }
