// /*
// * 1. Receive input from user from console until X key si pressed.package myassignment
// * 2. Received number is sorted and printed on the screen
// * 3. User is asked to enter the key until X key is pressed
//  */
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strconv"
// )

// func printSlice(s []int) {
// 	for _, value := range s {
// 		fmt.Println(value)
// 	}
// }

// func main() {
// 	var inValue string

// 	numList := []int{}
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for inValue != "X" {
// 		fmt.Print("Enter next number or X to exit: ")
// 		scanner.Scan()
// 		inValue = scanner.Text()
// 		i1, err := strconv.Atoi(inValue)
// 		if err != nil {
// 			//fmt.Println(err)
// 			break
// 		}
// 		numList = append(numList, i1)
// 		sort.Ints(numList)
// 		printSlice(numList)

// 	}

// }
