// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strconv"
// 	"strings"
// 	"sync"
// )

// func sortArray(arr *[]int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Goroutine is about to sort this subarray: %v\n", *arr)
// 	sort.Ints(*arr)
// }

// func main() {
// 	var sequence []int

// 	// Prompt the user for the sequence
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Type a space separated sequence of at least 4 integers: ")
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error processing the input:", err)
// 		return
// 	}

// 	// Trim off newline and whitespaces
// 	input = strings.TrimSpace(input)

// 	// Exit if trimmed input is empty
// 	if len(input) == 0 {
// 		fmt.Println("Empty input, exiting...")
// 		return
// 	}

// 	// Split and process into integers
// 	split_input := strings.Fields(input)
// 	sequence = make([]int, len(split_input))
// 	for i, v := range split_input {
// 		sequence[i], err = strconv.Atoi(v)
// 		if err != nil {
// 			fmt.Println("Error processing the input:", err)
// 			return
// 		}
// 	}

// 	// Exit if not enough integers
// 	if len(sequence) < 4 {
// 		fmt.Println("Not enough integers entered, exiting...")
// 		return
// 	}

// 	// Prepare a wait group
// 	var wg sync.WaitGroup
// 	wg.Add(4)

// 	// Split the sequence into 4 more or less equal parts
// 	// and start sorting goroutines immediately as a part is ready
// 	var arrays [4][]int
// 	step := int(len(sequence) / 4)
// 	for i := 0; i < 4; i++ {
// 		start := step * i
// 		end := start + step
// 		if i < 3 {
// 			arrays[i] = sequence[start:end]
// 		} else {
// 			arrays[i] = sequence[start:]
// 		}
// 		go sortArray(&arrays[i], &wg)
// 	}

// 	// Wait for all goroutines to finish executing
// 	wg.Wait()
// 	fmt.Printf("Individually sorted subarrays: %v\n", arrays)

// 	// Combine all 4 subarrays into the final array; do the final sort
// 	var finalArr []int
// 	for _, array := range arrays {
// 		finalArr = append(finalArr, array...)
// 	}
// 	sort.Ints(finalArr)

// 	fmt.Printf("Final sorted array: %v\n", finalArr)
// }
