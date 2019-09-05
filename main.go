package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandleBid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // TODO
}

func HandleJob(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // TODO
}

func main() {
	router := httprouter.New()

	router.GET("/bid/:jobid/:criteria", HandleBid)
	router.GET("/job/:criteria", HandleJob)

	log.Fatal(http.ListenAndServe(":8080", router))
}