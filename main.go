package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
)

func randomNumbers() int {
	return rand.Intn(1000)
}

// main - Output generates numbers randomly
func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("hihi3")
	fmt.Fprintf(w, "Random number is : %v", randomNumbers())
}

// go - console
func goHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "do next things")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/go", goHandler)

	// Bind to a port and pass our router in
	err := http.ListenAndServe(":8080", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}
