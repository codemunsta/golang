package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/ninja":
		fmt.Fprint(w, "Hello Caleb")
	default:
		fmt.Fprint(w, "Big fat error")
	}
	fmt.Printf("Handling function with %s request\n", r.Method)
}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Timeout Attempted")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func muskRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from new musk")
}
func main() {
	http.HandleFunc("/", htmlVsPlain)
	http.HandleFunc("/timeout", timeout)
	//http.ListenAndServe("", nil)

	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}
	var newMux http.ServeMux
	server.Handler = &newMux
	newMux.HandleFunc("/", muskRoute)
	server.ListenAndServe()
}
