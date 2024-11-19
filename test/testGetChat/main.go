package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// input token you got from responses and user login with whom you want to get chat
var (
	accessToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0IiwiZXhwIjoxNzM1NjI0MTUwfQ.aKBjfgNJ5LG9O4cEh0EoLstCD8gDj02lff4KjdLUrUc4_R793pOkdsUyADgxNQ3byFUDekChh7wr-2UXcxOXGwmsyEO8mRz_cGA56-AMWFRXPYzmY7ouW5kL77yHDY5g0yKzFc4_vohFETJz88aZnWFwsV8JuFHNrbaVPRX1hc3uM6DiGshQqpZquct6xB3I8nOIpMwwR_bQFmA_jDeBRMID9_e2euCRKKGWkNnzBqf6BurdDpRVgs0eKQsmMqFjJKp11y1mEQTNlKmmTU6NklSBrkUwRccU2twGT43xHIlhkg4j-sCWciBx4loANVO41qvBZnKKaFclPyJzPCOL-A"
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
