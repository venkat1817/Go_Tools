/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("commandsgit.txt")
	if err != nil {
		fmt.Println("error through thish page:", err)
		return
	}
	// defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
*/
