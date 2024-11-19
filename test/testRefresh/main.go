package main

import (
	"fmt"
	"log"
	"net/http"
)

// input tokens you got from responses(yes you need to do it manually)
var (
	accessToken  = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoxNzM1NjMzMTI4fQ.jmpLcOEl5splTNqoAr_RIqbpPOQRs1nntwWlyI6GFUQIDADG143smC76ftYvmd5f1-oGnnuRzG525mjs5YYp1VCIIiHlUaFXP4hSTEYtF5a76r-IRwCj3dcFYLLl7b3z76aWEZc9clz24fiUom80ItnRzqqG_AyDHh_N8mFdTLRnFg9V4t99d7Yu42wP3R_JOgCfe4_cGAOR6lDQ0aVCu2UHH6tL6qhUWRJBwTOmsml5-LO0Osv2ITgW7hdvDPQMZc0G8dMh3jbKuSyKCFGQuOqUxzdPkxwHKMeKlwb9yvcIjLTi0La-jksjL6TTTWzRya13NfTQRDfIBCRjc_0E_g"
	refreshToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzU2MzMxMjh9.bYzUO-bWidwvCsgwfOZML5MtusczOsRTHa76WqIQs4D_QCVceGEjaPDfWiWPH-ZZ40r3U9sGwjMfBylZX8wEDNBNpxzTpOJ2_YVSl8kKDaZ48DgAawwubzoV8qwryNSKuV9ALPPdCrLn_uW_TpE5z4vvWFSgU2QRYFtoIdGDwYQfQM1PfGYvtJTgIVIGbJXWLJB9M8y9vGnMalfa02oaLZY4Y3wlP25YVXUIPTiQm3crQepFiLBH1MLD3BrJhUOYFXf0ez7blr9He6AT8wv-KRE7P7wNvBzEgnAyAC0vnuDDAFS5A5xVSbkZ_2y4Dd-Eq8F-Fklok4BCGwYUG1qsVA"
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
	respAccessToken := resp.Header["X-Access-Token"] // tokens you get from responses(save it for later)🤓
	respRefreshToken := resp.Header["X-Refresh-Token"]
	fmt.Println(respAccessToken)
	fmt.Println(respRefreshToken)
}
