package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
)

func GetHelloEndpoint(w http.ResponseWriter, r *http.Request) {
	params := "Hello world"
	json.NewEncoder(w).Encode(params)
}

var responseMetric = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name: "request_duration_milliseconds",
		Help: "Request latency distribution",
		Buckets: prometheus.ExponentialBuckets(10.0, 1.13, 40),
	})

func main() {
	prometheus.MustRegister(responseMetric)
	router := mux.NewRouter()
	router.Handle("/metrics", prometheus.Handler())
	router.HandleFunc("/hello", GetHelloEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}