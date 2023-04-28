/*package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a Student struct to represent a student
type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

var students []Student // Slice of students to store the student data

func main() {
	// Define a handler function for the /students endpoint
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: // Handle GET request
			// Encode the slice of students as a JSON-encoded response
			json.NewEncoder(w).Encode(students)
		case http.MethodPost: // Handle POST request
			// Decode the request body into a Student struct
			var newStudent Student
			json.NewDecoder(r.Body).Decode(&newStudent)

			// Generate a unique ID for the new student
			newStudent.ID = len(students) + 1

			// Append the new student to the slice of students
			students = append(students, newStudent)

			// Encode the new student as a JSON-encoded response
			json.NewEncoder(w).Encode(newStudent)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Define a handler function for the /students/{id} endpoint
	http.HandleFunc("/students/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete { // Handle DELETE request
			// Get the student ID from the URL path parameter
			id := getIDFromURL(r.URL.Path)

			// Find the index of the student with the given ID in the slice of students
			index := findStudentIndexByID(id)

			if index == -1 { // Student not found
				http.NotFound(w, r)
				return
			}

			// Remove the student with the given ID from the slice of students
			students = append(students[:index], students[index+1:]...)

			// Return a success message
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Student deleted successfully"))
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the web server
	http.ListenAndServe(":6060", nil)
}

// Helper function to extract the student ID from the URL path parameter
func getIDFromURL(urlPath string) int {
	// The URL path should be in the form "/students/{id}"
	// Extract the ID as the last path segment
	var id int
	fmt.Sscanf(urlPath, "/students/%d", &id)
	return id
}

// Helper function to find the index of a student in the slice of students by ID
func findStudentIndexByID(id int) int {
	for i, student := range students {
		if student.ID == id {
			return i
		}
	}
	return -1 // Student not found
}
*/