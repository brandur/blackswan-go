package main

import (
	"fmt"
	"time"
)

type Event struct {
	Content    string    `json:"content"`
	OccurredAt time.Time `json:"occurred_at"`
	Slug       string    `json:"slug"`
	Type       string    `json:"type"`
}

func getEventsByType(event_type string) ([]*Event, error) {
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
