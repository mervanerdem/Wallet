package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	storage, _, err := NewMStorage("root:secret@tcp(127.0.0.1:3307)/Wallet")
	if err != nil {
		panic("Configration is wrong")
	}
	log.Fatal(http.ListenAndServe("localhost:8080", NewServer(storage)))
}
