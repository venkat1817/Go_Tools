// standed page display the success code in website for example http://example.com using go language.
// call the one file to onther file in websites.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("websites.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	websites := []string{}
	// Read website URLs from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		websiteURL := scanner.Text()
		websites = append(websites, websiteURL)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	for _, websiteURL := range websites {
		resp, err := http.Get(websiteURL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close() // Close the response body after reading

		fmt.Println("Service:", websites)
		fmt.Println("Status Code:", resp.StatusCode)

		// Check if the response status code indicates success (200-299)
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			fmt.Println("Success: HTTP request was successful")
		} else {
			fmt.Println("Error: HTTP request failed")
		}
		fmt.Println("-----------------------------")
	}
}
