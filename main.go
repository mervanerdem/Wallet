package main

import (
	"log"
	"net/http"
)

func main() {
	storage, _, err := NewMStorage(Path)
	if err != nil {
		panic("Configration is wrong")
	}
	log.Fatal(http.ListenAndServe(":8080", NewServer(storage)))
}
