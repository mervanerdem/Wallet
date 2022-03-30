package main

import (
	"log"
	"net/http"
)

func main() {
	storage, _, _ := NewMStorage("'root':Mervan.1907@tcp(127.0.0.1:3306)/wallet_db?charset=utf8mb4&parseTime=True&loc=Local")
	log.Fatal(http.ListenAndServe(":8080", NewServer(storage)))
}
