package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Another update\n")
		if err != nil {
			panic(err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
