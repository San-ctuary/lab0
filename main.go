package main

import (
	"log"
	"main/handlers"
	"net/http"
)

func main() {
	// TODO: some code goes here
	// Fill out the HomeHandler function in handlers/handlers.go which handles the user's GET request.
	// Start an http server using http.ListenAndServe that handles requests using HomeHandler.
	http.HandleFunc("/", handlers.HomeHandler)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:	", err)
	}

	// test
	//db := ridershipDB.CsvRidershipDB{}
	//db.Open("mbta.csv")
	//values, _ := db.GetRidership("blue")
	//fmt.Printf("%v", values)
}
