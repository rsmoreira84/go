package main

/*
Simple web server

To run it execute:
	$ go run main.go &

	Then open a browser and go to:
		http://localhost:8000/
		http://localhost:8000/count

To stop the server, use the `kill` command with the process ID of the running server:

	$ lsof -i :8000
	$ kill <pid>

*/

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/rsmoreira84/go/lissajous"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/gif", lissajousgif)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajousgif(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	lissajous.LissajousCycles(w, 20)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Counters = %d\n", count)
	mu.Unlock()
}
