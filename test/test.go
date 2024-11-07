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
	body, _ := json.Marshal(models.User{Login: "login", Password: "password"})
	r := bytes.NewReader(body)
	request, err := http.NewRequest("POST", "http://localhost:8082/register", r)
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("dadada")
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
}
