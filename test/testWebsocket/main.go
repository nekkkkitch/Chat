package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/coder/websocket"
)

type message struct {
	Reciever string `json:"reciever"`
	Text     string `json:"text"`
}

// input token you got from responses and user login who you want to send message to
var (
	accessToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM1NjMzMzk2fQ.b80929QbkVIyjbwmXOky7Mn-eTB33FftIq7nV53QH4et8KT_HG5BVbjzbJ8xiTOpHRZV2uFPf4hoZkl9eCR6nBF5rYOuNr2Fe83plCdjiA0Jd0g2W5g6Au3_E5YMnxUtOu4JUp5-5GmfSjVttWF-x_T-cknw7bjUhk_iQGOmKhlbDhrgYk3tQHfd7t5ncEPcScAkhjrhVatF6D_uO1zstU0GsaF2sp4HdedPiKMx6zXuVqFwMMkMmYU_rZUvE1lNWYDH6C9bTdE0fWTfcD9ag8lEClpiyRW8vZJrTE6A8ns2sGcRhJTfngQg8iDOGuu3UQBcQhPfTReYdtL854mWpw"
	resp        = make(chan string)
)

func main() {
	c, _, err := websocket.Dial(context.Background(), fmt.Sprintf("ws://localhost:8082/chat?access_token=%v", accessToken), nil)
	if err != nil {
		log.Fatalln(err)
	}
	message := message{Reciever: "login", Text: "test"} // feel free to change something here
	msg, _ := json.Marshal(message)
	c.Write(context.Background(), websocket.MessageBinary, msg) //send message to user(it doesnt matter if he is online or not)
	//waiting for messages
	go waitForMessage(c)
	for {
		msg := <-resp
		log.Println("\n" + msg)
	}
}

func waitForMessage(c *websocket.Conn) { // message reader(you probably want to make another user for this to work or just send messages to yourself)
	for {
		_, msgBytes, err := c.Read(context.Background())
		resp <- string(msgBytes)
		if err != nil {
			log.Println(err)
		}
	}
}
