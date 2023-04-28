//https://medium.com/hackernoon/deploying-a-go-application-on-aws-ec2-76390c09c2c5

// Thise link is used to deploy a Go application on AWS ec2
package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello worl ")
}

func main() {
	http.HandleFunc("/venkat", HomeEndpoint)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
