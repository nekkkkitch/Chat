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
	MessageText string `json:"message_text"`
	Sender      int    `json:"message_sender"`
	Reciever    string `json:"message_reciever"`
}

type BeautifiedMessage struct {
	MessageText string `json:"message_text"`
	Sender      int    `json:"message_sender"`
	Reciever    int    `json:"message_reciever"`
}
