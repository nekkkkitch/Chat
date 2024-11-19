package main

import (
	"Chat/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	body, _ := json.Marshal(models.User{Login: "logiiiiin", Password: "password"})
	r := bytes.NewReader(body)
	request, err := http.NewRequest("POST", "http://localhost:8082/login", r)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	accessToken := resp.Header["X-Access-Token"]
	refreshToken := resp.Header["X-Refresh-Token"]
	fmt.Println(accessToken)
	fmt.Println(refreshToken)
}
