// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func BubbleSort(numbers []int) {
// 	n := len(numbers)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if numbers[j] > numbers[j+1] {
// 				Swap(numbers, j)
// 			}
// 		}
// 	}
// }

// func Swap(numbers []int, i int) {
// 	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
// }

// func main() {
// 	var input string
// 	var numbers []int

// 	fmt.Println("Enter up to 10 integers (press X to stop):")
// 	for i := 0; i < 10; i++ {
// 		fmt.Scan(&input)
// 		if input == "X" {
// 			break
// 		}
// 		num, err := strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Invalid input, Enter X to stop")
// 			continue
// 		}
// 		numbers = append(numbers, num)
// 	}

// 	fmt.Println("Unsorted list of integers:", numbers)

// 	BubbleSort(numbers)

// 	fmt.Println("Sorted list of integers:", numbers)
// }
