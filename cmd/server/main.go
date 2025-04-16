package main

import (
	"log"

	"github.com/gabeno/logstore/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8000")
	log.Fatal(srv.ListenAndServe())
}
