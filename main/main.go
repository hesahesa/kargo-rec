package main

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"github.com/hesahesa/kargorec/datasource"
	"github.com/hesahesa/kargorec"
)

var (
	FakeDataSource = datasource.NewFakeDataSource()
	Kargo = kargorec.NewKargo(FakeDataSource)
)

func HandleBid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	jobID, _ := strconv.Atoi(p.ByName("jobid"))
	criteria := p.ByName("criteria")
	bids, _ := Kargo.BidbyCriteria(jobID, criteria)

	b, err := json.Marshal(bids)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return;
	}
	
	fmt.Fprint(w, string(b))
}

func HandleJob(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	criteria := p.ByName("criteria")
	jobs, _ := Kargo.JobbyCriteria(criteria)

	b, err := json.Marshal(jobs)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return;
	}
	
	fmt.Fprint(w, string(b))
}

func main() {
	router := httprouter.New()

	router.GET("/bid/:jobid/:criteria", HandleBid)
	router.GET("/job/:criteria", HandleJob)

	log.Fatal(http.ListenAndServe(":8080", router))
}
