package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

const (
	baseURL  = "https://discord.com/api/v9/users/@me/pomelo-attempt"
	token    = "user-token"
	userName = "username"
)

type Response struct {
	Taken bool `json:"taken"`
}

func main() {

	if len(userName) < 2 || len(userName) > 32 {
		fmt.Println("Usernames must be at least 2 characters and at most 32 characters long")
		return
	}

	if !validUsername(userName) {
		fmt.Println("Invalid username format. Permitted characters for new usernames: \nLatin characters (a-z) \nNumbers (0-9) \nUnderscore ( _ )  \nPeriod ( . )")
		return
	}

	Request(userName)
}

func validUsername(userName string) bool {
	regex := "^[a-z0-9_.]+$"
	match, _ := regexp.MatchString(regex, userName)
	return match
}

func Request(userName string) {
	client := &http.Client{}

	body, _ := json.Marshal(map[string]string{
		"username": userName,
	})

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
	}

	req.Header = http.Header{
		"Authorization": {token},
		"Content-Type":  {"application/json"},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var status Response
	err = json.Unmarshal(body, &status)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(Message(status.Taken, userName))
}

func Message(taken bool, userName string) string {
	switch taken {
	case true:
		return userName + " is not available."
	default:
		return userName + " is available!"
	}
}
