// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// )

// func main() {
// 	// make slice with capacity of 3
// 	slice := make([]int, 0, 3)

// 	for {
// 		var input string
// 		fmt.Println("Enter an integer to add to the slice (enter 'X' to exit):")
// 		_, err := fmt.Scanln(&input)
// 		if err != nil {
// 			fmt.Println("Error reading input:", err)
// 			return
// 		}

// 		if input == "X" {
// 			break
// 		}

// 		// input string to integer
// 		num, err := strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Please enter a valid integer or 'X' to exit.")
// 			continue
// 		}

// 		slice = append(slice, num)
// 		sort.Ints(slice)
// 		// print sorted slice
// 		fmt.Println("Sorted slice:", slice)
// 	}
// }
