// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// // Define a name struct
// type Name struct {
// 	Fname string
// 	Lname string
// }

// func main() {
// 	var filename string

// 	// Prompt user for the name of the text file
// 	fmt.Println("Enter the name of the text file:")
// 	_, err := fmt.Scanln(&filename)
// 	if err != nil {
// 		fmt.Println("ERROR reading filename:", err)
// 		return
// 	}

// 	// Open the file
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println("ERROR opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Create a slice of structs to hold names
// 	var names []Name

// 	// Read lines from the file and create structs
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		// split string into individual words based on spaces
// 		fields := strings.Fields(line)

// 		// Ensure that the line has both first and last names
// 		if len(fields) >= 2 {
// 			fname := fields[0]
// 			lname := fields[1]

// 			// Create a Name struct and add it to the slice
// 			name := Name{fname, lname}
// 			names = append(names, name)
// 		}
// 	}

// 	// Check for any scanning error
// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("ERROR scanning file:", err)
// 		return
// 	}

// 	// Iterate through the slice of structs and print names
// 	for _, name := range names {
// 		fmt.Printf("First Name: %-20s Last Name: %-20s\n", name.Fname, name.Lname)
// 	}
// }
