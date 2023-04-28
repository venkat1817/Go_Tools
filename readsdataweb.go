package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/venkat", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("commandsgit.txt")
		if err != nil {
			http.Error(w, "Error opening file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Fprintln(w, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			http.Error(w, "Error scanning file", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8081", nil)
}
