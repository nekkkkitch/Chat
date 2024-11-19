package models

import "time"

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Message struct {
	Text     string    `json:"text"`
	Sender   int       `json:"sender"`
	Reciever string    `json:"reciever"`
	SendTime time.Time `json:"time"`
	Hash     string    `json:"hash"`
}

type BeautifiedMessage struct {
	Text     string `json:"text"`
	Sender   int    `json:"sender"`
	Reciever int    `json:"reciever"`
	Hash     string `json:"hash"`
}
