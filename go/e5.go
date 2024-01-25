// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	var name, address string

// 	// Prompt user to enter name
// 	fmt.Println("Enter a name:")
// 	_, err := fmt.Scanln(&name)
// 	if err != nil {
// 		fmt.Println("ERROR reading input name:", err)
// 		return
// 	}

// 	// Prompt user to enter address
// 	fmt.Println("Enter an address:")
// 	_, err = fmt.Scanln(&address)
// 	if err != nil {
// 		fmt.Println("ERROR reading input address:", err)
// 		return
// 	}

// 	// Create a map and add name and address
// 	data := map[string]string{
// 		"name":    name,
// 		"address": address,
// 	}

// 	// Marshal the map to JSON
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println("ERROR marshaling JSON:", err)
// 		return
// 	}

// 	// Print the JSON object
// 	fmt.Println("JSON object:")
// 	fmt.Println(string(jsonData))
// }
