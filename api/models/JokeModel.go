package Models

type Joke struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
	Joke  string `json:"joke"`
}

func (b *Joke) TableName() string {
	return "joke"
}
