package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO! PATH: %s", r.URL.Path[1:])
}

func handleBuy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Buying Stuff")
}

func main() {
	http.HandleFunc("/", handleAll)
	http.HandleFunc("/buy", handleBuy)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
