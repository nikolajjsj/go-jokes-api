package Models

import "time"

type Joke struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Joke      string    `json:"joke"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Joke) TableName() string {
	return "jokes"
}
