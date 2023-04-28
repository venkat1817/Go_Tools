// web application display the success code in website for example http://example.com using go language
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Open the websites file
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Loop through each website URL and check their success codes
		for _, websiteURL := range websites {
			resp, err := http.Get(websiteURL)
			if err != nil {
				// Display the website URL and error message in the HTTP response
				fmt.Fprintf(w, "Website: %s\nError: %s\n\n", websiteURL, err)
				continue
			}
			// defer resp.Body.Close()

			// Display the website URL and its success code in the HTTP response
			fmt.Fprintf(w, "Website: %s\nStatus: %s\n\n", websiteURL, resp.Status)
		}
	})

	http.ListenAndServe(":8080", nil)
}
