package main

import (
	"log"
	"os"
)

// Get device host name
func getDevice() string {
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return host
}
