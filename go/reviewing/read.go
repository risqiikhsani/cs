// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strings"
// )

// type People struct {
// 	fname string
// 	lname string
// }

// func main() {
// 	fmt.Println("enter a file name:")
// 	in := bufio.NewReader(os.Stdin)
// 	line, err := in.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error while inserting file name. Please reload the program")
// 	}
// 	line = strings.TrimSuffix(line, "\n")
// 	content, err := ioutil.ReadFile(line)
// 	if err != nil {
// 		fmt.Println("Cant read file \n Aborting.... \n Restart the program")
// 		os.Exit(0)
// 	}
// 	var contentAsString = strings.Split(string(content),"\n")
// 	var users = make([]People, 0)
// 	for i := range contentAsString {
// 		users = append(users, convertStringToPeople(contentAsString[i]))
// 	}

// 	for i := range users {
// 		fmt.Println("user : " + users[i].fname + " " + users[i].lname)
// 	}

// }

// func convertStringToPeople(input string) People {
// 	var people = People{}
// 	var parsedInput = strings.Split(input," ")
// 	people.fname = parsedInput[0]
// 	people.lname = parsedInput[1]
// 	return people
// }
