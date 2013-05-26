package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "time"
)

type Event struct {
	Content    string    `json:"content"`
	OccurredAt time.Time `json:"occurred_at"`
	Slug       string    `json:"slug"`
	Type       string    `json:"type"`
}

func main() {
    http.HandleFunc("/events", events)
    fmt.Println("listening")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func queryEvents(event_type string) ([]*Event, error) {
	rows, err := pg.Query(`
        select content, occurred_at, slug, type
        from events
        where type = $1
        order by occurred_at
        `,
        event_type)
	defer rows.Close()
	var events []*Event
	if err != nil {
        fmt.Println("error=%s", err)
		return nil, err
	}
	for rows.Next() {
		event := new(Event)
		rows.Scan(&event.Content, &event.OccurredAt, &event.Slug, &event.Type)
		events = append(events, event)
    }
    return events, nil
}

func events(res http.ResponseWriter, req *http.Request) {
	startResponse := time.Now()
    events, err := queryEvents("blog")
	if err != nil {
		fmt.Printf("error=\"unable to fetch events\"\n")
		os.Exit(1)
	}
    encoded, err := json.Marshal(events)
    if err != nil {
      panic(err)
    }
    fmt.Fprintln(res, string(encoded))
	MeasureT(startResponse, "instrumentation.events")
}
