package models

type News struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Text    string `json:"text"`
	Picture string `json:"picture"`
}
