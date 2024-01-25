// package main

// import (
// 	"fmt"
// 	"sort"
// 	"sync"
// )

// func sortPartition(wg *sync.WaitGroup, arr []int, index int) {
// 	defer wg.Done()
// 	fmt.Println("=====================================")
// 	fmt.Println("index of : ", index)
// 	partitionSize := len(arr) / 4
// 	fmt.Println("partition size : ", partitionSize)
// 	start := index * partitionSize
// 	fmt.Println("start from : ", start)
// 	end := start + partitionSize
// 	fmt.Println("end at : ", end)

// 	if index == 3 {
// 		end = len(arr)
// 	}

// 	part := arr[start:end]
// 	fmt.Printf("Partition %d: %v\n", index+1, part)

// 	// sort the sub array
// 	sort.Ints(part)
// 	fmt.Printf("Sorted Partition %d: %v\n", index+1, part)
// 	fmt.Println("==============================")
// 	// copy sorted array to original array
// 	copy(arr[start:end], part)
// }

// func merge(arr []int) {
// 	sort.Ints(arr)
// 	fmt.Println("Entire sorted list : ", arr)
// }

// func main() {
// 	var input int
// 	var numbers []int
// 	fmt.Println("Enter number (enter non number to stop) : ")
// 	for {
// 		_, err := fmt.Scan(&input)
// 		if err != nil {
// 			break
// 		}
// 		numbers = append(numbers, input)
// 	}

// 	if len(numbers) < 4 {
// 		fmt.Println("Please enter atleast 4 numbers")
// 		return
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(4)

// 	for i := 0; i < 4; i++ {
// 		go sortPartition(&wg, numbers, i)
// 	}

// 	wg.Wait()
// 	merge(numbers)
// }
