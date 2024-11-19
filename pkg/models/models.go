package models

import "time"

// Данные о пользователе
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Данные для аутентификации
type AuthData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Данные о полученном от пользователя сообщении
type Message struct {
	Text     string    `json:"text"`
	Sender   int       `json:"sender"`
	Reciever string    `json:"reciever"`
	SendTime time.Time `json:"time"`
	Hash     string    `json:"hash"`
}

// Данные о сообщении, возвращаемого пользователю
type BeautifiedMessage struct {
	Text     string `json:"text"`
	Sender   int    `json:"sender"`
	Reciever int    `json:"reciever"`
	Hash     string `json:"hash"`
}
