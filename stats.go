package covid19

import "time"

type CoronaStats struct {
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	Confirmed int64     `json:"confirmed"`
	Recovered int64     `json:"recovered"`
	Deaths    int64     `json:"deaths"`
}

type News struct {
	ID     string    `json:"id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
	Body   string    `json:"body"`
	Media  []string  `json:"media"`
	Source string    `json:"source"`
}
