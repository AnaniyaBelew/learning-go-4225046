package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reading user input with type Conversion
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please Enter Your name: ")
	str, _ := reader.ReadString('\n')
	fmt.Printf("Hello Mr/Ms %s\n", str)
	fmt.Print("Please Enter Your Age: ")
	age, _ := reader.ReadString('\n')
	parsedAge, err := strconv.ParseInt(strings.TrimSpace(age), 0, 64)
	if err != nil {
		fmt.Print("Error Parsing Your Age", err)
	} else {
		fmt.Printf("Great Your Age is: %d\n", parsedAge)
	}
}
