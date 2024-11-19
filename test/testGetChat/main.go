package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// input token you got from responses and user login with whom you want to get chat
var (
	accessToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM1NjMzMzk2fQ.b80929QbkVIyjbwmXOky7Mn-eTB33FftIq7nV53QH4et8KT_HG5BVbjzbJ8xiTOpHRZV2uFPf4hoZkl9eCR6nBF5rYOuNr2Fe83plCdjiA0Jd0g2W5g6Au3_E5YMnxUtOu4JUp5-5GmfSjVttWF-x_T-cknw7bjUhk_iQGOmKhlbDhrgYk3tQHfd7t5ncEPcScAkhjrhVatF6D_uO1zstU0GsaF2sp4HdedPiKMx6zXuVqFwMMkMmYU_rZUvE1lNWYDH6C9bTdE0fWTfcD9ag8lEClpiyRW8vZJrTE6A8ns2sGcRhJTfngQg8iDOGuu3UQBcQhPfTReYdtL854mWpw"
	user        = "login"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8082/getchat?user=%v", user), nil)
	request.Header["X-Access-Token"] = append(request.Header["X-Access-Token"], accessToken)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bodyBytes))
}
