package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Discord struct {
	hook string
}

type discordPayload struct {
	Content string `json:"content"`
}

func NewDiscord(webhook string) Discord {
	return Discord{
		hook: webhook,
	}
}

func (d *Discord) push_message(msg string) {
	payload := discordPayload{
		Content: msg,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		panic("cannot marshal message")
	}

	req, err := http.NewRequest("POST", d.hook, bytes.NewBuffer(data))
	if err != nil {
		panic("cannot send request to discord")
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Errorf("error sending request to discord %v", err))
	}
	defer resp.Body.Close()
}
