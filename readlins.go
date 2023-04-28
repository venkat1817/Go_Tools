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
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; i < 4; i++ {
		if scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
	for i := 0; i < 3; i++ {
		if scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
*/
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
		fmt.Println("Error opening file:", err)
		return
	}
	scanner := bufio.NewScanner(file)
    for i := 1; i <= 10; i++ {
		if !scanner.Scan() {
			fmt.Println("Error: file has less than 10 lines")
			return
		}
		if i == 9 {
			fmt.Println(scanner.Text())
			break
		}
	}
    if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}
*/
