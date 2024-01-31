package main

import (
	"log"
	"net/http"

	"github.com/rusher2004/go-rest-api/datastore"
	"github.com/rusher2004/go-rest-api/db"
	"github.com/rusher2004/go-rest-api/server"

	_ "github.com/lib/pq"
)

func main() {
	db, err := db.NewDB("postgres", "some dsn")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	d := datastore.NewDataStore("some api client", &db)
	s := server.NewServer(d)

	log.Println("starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
