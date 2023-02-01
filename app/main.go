package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/saya", endointsaya)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {

		panic(err)

	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World Lagi!")
}

func endointsaya(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ini adalah endpoint saya!")
}
