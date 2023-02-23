package main

import (
	"fmt"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World putra!!")
}

func Endointsaya(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ini adalah endpoint saya!")
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/saya", Endointsaya)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {

		panic(err)

	}
}
