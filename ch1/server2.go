package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var cnt int64

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/count", count)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	cnt++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Patch=%s\n", r.URL.Path)
}

func count(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count=%d\n", cnt)
	mu.Unlock()
}
