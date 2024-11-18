package models

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
	MessageText string `json:"text"`
	Sender      int    `json:"sender"`
	Reciever    string `json:"reciever"`
	Hash        string `json:"hash"`
}

type BeautifiedMessage struct {
	MessageText string `json:"text"`
	Sender      int    `json:"sender"`
	Reciever    int    `json:"reciever"`
	Hash        string `json:"hash"`
}
