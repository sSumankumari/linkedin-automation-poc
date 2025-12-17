package types

import "time"

type Profile struct {
	URL       string
	Collected time.Time
	Processed bool
}

type Connection struct {
	ProfileURL string
	SentAt     time.Time
	Status     string
}

type Message struct {
	ProfileURL string
	Content    string
	SentAt     time.Time
}
