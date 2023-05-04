package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	baseURL  = "https://discord.com/api/v10/users/@me/pomelo"
	token    = "user-token"
	userName = "username"
)

type Response struct {
	Code int `json:"code"`
}

func main() {
	client := &http.Client{}

	body, _ := json.Marshal(map[string]string{
		"username": userName,
	})

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("%v", err)
	}

	req.Header = http.Header{
		"Authorization": {token},
		"Content-Type":  {"application/json"},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var status Response
	err = json.Unmarshal(body, &status)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println(Message(status.Code, userName))
}

func Message(code int, userName string) string {
	switch code {
	case 40001:
		return userName + " is available!"
	case 50035:
		return userName + " is not available."
	default:
		return "Unknown error."
	}
}
