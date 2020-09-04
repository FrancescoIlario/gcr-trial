package main

import (
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello!"))
	})

	err := http.ListenAndServe(":8080", nil)
	return err
}
