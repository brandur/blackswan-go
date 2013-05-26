package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/events", events)
	fmt.Println("listening")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func events(res http.ResponseWriter, req *http.Request) {
	startResponse := time.Now()

	events, err := getEventsByType("blog")
	if err != nil {
		fmt.Printf("error=\"unable to fetch events\"\n")
		os.Exit(1)
	}
	MeasureT(startResponse, "get.events")

	startMarshal := time.Now()
	encoded, err := json.Marshal(events)
	if err != nil {
		panic(err)
	}
	MeasureT(startMarshal, "marshal.events")

	fmt.Fprintln(res, string(encoded))
	MeasureT(startResponse, "instrumentation.events")
}
