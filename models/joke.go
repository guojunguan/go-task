package models

type JokeValue struct {
	Id         uint     `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

type JokeResponse struct {
	JokeType  string    `json:"type"`
	Last_name string    `json:"last_name"`
	Value     JokeValue `json:"value"`
}
