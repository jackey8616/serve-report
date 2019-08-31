package main

import (
	"log"
	"net/http"

	"sereport/sereport"

	"gopkg.in/mgo.v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dataPath := "./data"

	session, _ := mgo.Dial("mongo:27017")
	db := session.DB("serve-report")

	sr := sereport.NewSereport(db, &dataPath)
	router := sr.RegisterRoute()

	fs := http.FileServer(http.Dir("./data/repo/"))
	http.Handle("/repo/", http.StripPrefix("/repo/", fs))

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}
