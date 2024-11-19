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

// input token you got from responses and user login with whom you want to get chat
var (
	accessToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0IiwiZXhwIjoxNzM1NjI0MTUwfQ.aKBjfgNJ5LG9O4cEh0EoLstCD8gDj02lff4KjdLUrUc4_R793pOkdsUyADgxNQ3byFUDekChh7wr-2UXcxOXGwmsyEO8mRz_cGA56-AMWFRXPYzmY7ouW5kL77yHDY5g0yKzFc4_vohFETJz88aZnWFwsV8JuFHNrbaVPRX1hc3uM6DiGshQqpZquct6xB3I8nOIpMwwR_bQFmA_jDeBRMID9_e2euCRKKGWkNnzBqf6BurdDpRVgs0eKQsmMqFjJKp11y1mEQTNlKmmTU6NklSBrkUwRccU2twGT43xHIlhkg4j-sCWciBx4loANVO41qvBZnKKaFclPyJzPCOL-A"
	resp        = make(chan string)
)

func main() {
	c, _, err := websocket.Dial(context.Background(), fmt.Sprintf("ws://localhost:8082/chat?access_token=%v", accessToken), nil)
	if err != nil {
		log.Fatalln(err)
	}
	message := message{Reciever: "logsain", Text: "test"} // feel free to change something here
	msg, _ := json.Marshal(message)
	c.Write(context.Background(), websocket.MessageBinary, msg)
	//waiting for messages
	go waitForMessage(c)
	for {
		msg := <-resp
		log.Println(msg)
	}
}

func waitForMessage(c *websocket.Conn) {
	for {
		_, msgBytes, err := c.Read(context.Background())
		resp <- string(msgBytes)
		if err != nil {
			log.Println(err)
		}
	}
}
