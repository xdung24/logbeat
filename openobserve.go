package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Pusher struct {
	Email    string
	Password string
	Host     string
	Path     string
	Client   *http.Client
}

func (p *Pusher) pushLog(file string, timestamp string, content string, device string, port int) {
	startTime := time.Now() // Record the start time

	data := fmt.Sprintf(`{"file":"%s", "time": "%s", "content": "%s", "port": %d, "device": "%s"}`, file, timestamp, content, port, device)
	req, err := http.NewRequest("POST", fmt.Sprintf(`%s%s`, p.Host, p.Path), strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(p.Email, p.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	endTime := time.Now()              // Record the end time
	duration := endTime.Sub(startTime) // Calculate the duration
	log.Printf("%s (%s)\n", string(body), duration)
}
