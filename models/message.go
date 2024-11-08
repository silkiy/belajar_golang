package models

type Message struct {
	ID      int    `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
