package main

import (
	"fmt"
	"log"
	"net/http"
)

// input tokens you got from responses
var (
	accessToken  = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI1IiwiZXhwIjoxNzM1NjI3MzQyfQ.JU2rx0WANcH3LSs964n-oI6S21HksnjOg5MFucfcXfmW2soyNXvcgC1aGDU0vNKKNd1WQdb1z5Ji2tIMrW0IMHCwvTvtlySoX4KOCHJdfj4222UuFLkfLKfMCTr6C4fG36blyMTvYmGUaOtfoPNoKlVDuomudvGcpk1ak_ujzQVyha5RN1-CilgZNDa5Mw7rKplV7HS3I8B4mUyEbe8NsWXvbgFjh2Z9CmZk-aRuD9-fdh6PGLub9AOIKcQGK6sP1lovA4n7B35ikazvlY7hdAqGCG1SB16ErWCiwjjfU9eHjMeDAGJ7iuLfI2G4MbVY9dtNHWiM81paUqyPswXP0Q"
	refreshToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzU2MjczNDJ9.D7GNHg03doI4cG3tUmCq8C17uRHfDGdUD1tuXQhToE0099f_0OKb0FpnTdouNHEtShBv2HkpuPviSzHt0TXX_SqjUdGva_kZw73pMKiAFaK0SU9GelqbbI6KZBkEJ8DVLStmFNlB4Dz6VbdmfeqqCJW526HPYZ7P2HWoCHwPwu614iAQbJiDlm8PutpQ6XxXKeFSzVesbkY_-Q3AgpfcBejqmyMrzYiK1c-pxruEve88wdKxoWPD8vgakBt1Kxzw6cDQiMH0hszv1kBkV_uPm8fAnJw6cWU1YovaS_tkCpwmBB6SPfm2o5xBi1Nwmh9iJpsUm_HTJvIZhD8ni0tPPQ"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8082/refresh", nil)
	request.Header["X-Access-Token"] = append(request.Header["X-Access-Token"], accessToken)
	request.Header["X-Refresh-Token"] = append(request.Header["X-Refresh-Token"], refreshToken)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	respAccessToken := resp.Header["X-Access-Token"]
	respRefreshToken := resp.Header["X-Refresh-Token"]
	fmt.Println(respAccessToken)
	fmt.Println(respRefreshToken)
}
