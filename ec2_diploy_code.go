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

// package main

// import (
// 	"fmt"
// 	_ "io/ioutil"
// 	"log"
// 	"net/http"

// 	"golang.org/x/crypto/ssh"
// )

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// Establish an SSH connection to the EC2 instance
// 	config := &ssh.ClientConfig{
// 		User: "venkat",
// 		Auth: []ssh.AuthMethod{
// 			ssh.Password("<your EC2 instance password>"),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 	}
// 	conn, err := ssh.Dial("tcp", "52.77.249.216:22", config)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer conn.Close()

// 	// Execute a command on the remote server
// 	session, err := conn.NewSession()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer session.Close()
// 	output, err := session.Output("echo 'Hello World!'")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Render the output in the browser
// 	fmt.Fprintf(w, "<html><body><pre>%s</pre></body></html>", string(output))
// }
