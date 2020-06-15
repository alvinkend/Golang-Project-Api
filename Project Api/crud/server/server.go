package main

import (
	"log"
	"net/http"

	"github.com/bootcampProject/crud"
)

var (
	dbUser    = "postgres"
	dbPass    = "postgres"
	dbDefault = "postgres"
	dbName    = "databaseboongan"
)

func main() {
	db, err := crud.Connect(dbUser, dbPass, dbName)
	if err != nil {
		log.Fatal(err)
	}
	crud.RegisDB(db)
	http.HandleFunc("/api/ss/", crud.SS)
	defer db.Close()

	log.Println("localhost : 8080")
	http.ListenAndServe(":8080", nil)
}
