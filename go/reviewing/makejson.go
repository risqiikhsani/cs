// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// func main() {

// 	inputReader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Please enter a name:")
// 	inputName, _ := inputReader.ReadString('\n')
// 	inputName = inputName[:len(inputName)-1]
// 	fmt.Println("Please enter an address:")
// 	inputAddress, _ := inputReader.ReadString('\n')
// 	inputAddress = inputAddress[:len(inputAddress)-1]

// 	idMap := map[string]string{
// 		"name":    inputName,
// 		"address": inputAddress}
// 	fmt.Println(idMap)

// 	barr, _ := json.Marshal(idMap)
// 	fmt.Println(string(barr))

// }
