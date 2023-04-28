/*
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> iam  venkat !</h1>")

}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
*/
/*
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to read file: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", html)
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
*/