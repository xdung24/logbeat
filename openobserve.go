package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Pusher struct {
	Email    string
	Password string
	Host     string
	Path     string
	Timeout  int
}

func (p *Pusher) pushLog(file string, time string, content string, device string, port int) {
	data := fmt.Sprintf(`{"file":"%s", "time": "%s", "content": "%s", "port": %d, "device": "%s"}`, file, time, content, port, device)
	req, err := http.NewRequest("POST", fmt.Sprintf(`%s%s`, p.Host, p.Path), strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(p.Email, p.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}
