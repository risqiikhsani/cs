// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// )

// func divideSliceIntoFour(s []int) [][]int {
// 	var dividedSlice [][]int

// 	size := len(s) / 3

// 	if size < 1 {
// 		for _, value := range s {
// 			dividedSlice = append(dividedSlice, []int{value})
// 		}
// 		return dividedSlice
// 	}

// 	for i := 0; i < len(s); i += size {
// 		split := i + size

// 		if split > len(s) {
// 			split = len(s)
// 		}

// 		dividedSlice = append(dividedSlice, s[i:split])
// 	}

// 	return dividedSlice
// }

// func sortSubslice(s []int, ch chan []int) {
// 	fmt.Println("Subslice before sorting: ", s)
// 	sort.Ints(s)
// 	ch <- s
// }

// func main() {
// 	var unsorted []int
// 	for {
// 		var value string
// 		fmt.Println("Please enter an integer. When you are ready to sort, enter 'sort'.")
// 		fmt.Print("> ")

// 		fmt.Scan(&value)
// 		if value == "sort" {
// 			break
// 		}
// 		intValue, err := strconv.Atoi(value)
// 		if err != nil {
// 			fmt.Println("Value != sort and could not parse to int")
// 		}
// 		unsorted = append(unsorted, intValue)
// 	}
// 	fmt.Println("Array to sort: ", unsorted)

// 	var sorted []int
// 	c := make(chan []int)

// 	for _, subSlice := range divideSliceIntoFour(unsorted) {
// 		go sortSubslice(subSlice, c)
// 	}

// 	for {
// 		select {
// 		case s := <-c:
// 			sorted = append(sorted, s...)
// 			if len(sorted) == len(unsorted) {
// 				sort.Ints(sorted)
// 				fmt.Println("Sorted array: ", sorted)
// 				return
// 			}
// 		}
// 	}
// }
