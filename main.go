package main

import (
	"log"
	"net/http"
	"wallet/mysql"
	"wallet/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	storage, _, err := mysql.NewMStorage("tester:secret@tcp(db:3306)/test")
	if err != nil {
		panic("Configration is wrong")
	}
	log.Fatal(http.ListenAndServe("localhost:8080", server.NewServer(storage)))
}
