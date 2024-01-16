package main

import (
	"log"
	"net/http"

	newstore "github.com/rusher2004/go-rest-api/new-store"
	oldstore "github.com/rusher2004/go-rest-api/old-store"
	"github.com/rusher2004/go-rest-api/server"
)

func main() {
	new := newstore.NewDataStore("some api client", "some db connection")
	old := oldstore.NewDataStore("some lambda client")
	s := server.NewServer(new, old)

	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
