package adapter

import (
	"fmt"
	"log"
	"net/http"
)

type MyServer struct {
	Msg string
}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func run() {
	// server := &MyServer{
	// 	Msg: "Hello, World",
	// }

	// http.Handle("/", server)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	log.Fatal(http.ListenAndServe(":10080", nil))
}
