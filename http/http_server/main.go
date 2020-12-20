package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		sayBye(w, r)
	})
	server := &http.Server{
		Addr:         ":8081",
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	log.Println("Starting httpserver at 8080")
	log.Fatal(server.ListenAndServe())
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Write([]byte("bye bye,http server"))
}
