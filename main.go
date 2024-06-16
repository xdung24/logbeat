package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/radovskyb/watcher"
)

var password string

func main() {
	// Define a flag for the folder path
	var folderPath string
	flag.StringVar(&folderPath, "folder", `./logs/`, "Path to the folder to watch")

	// Parse the flags
	flag.Parse()

	// Verify the folder path exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		fmt.Println("No folder path provided. Usage:")
		flag.Usage()
		os.Exit(1)
	}

	// Get the device name
	var deviceName = getDevice()

	// Create a new file watcher.
	w := watcher.New()
	w.SetMaxEvents(1)          // Only allow one event to be processed at a time
	w.FilterOps(watcher.Write) // Only watch for write events

	// Watch a specific folder recursively for changes.
	if err := w.AddRecursive(folderPath); err != nil {
		log.Fatalln(err)
	}

	// Create a custom HTTP client with a timeout and keep-alive settings
	customClient := &http.Client{
		Timeout: time.Duration(15) * time.Second, // Set the timeout
		Transport: &http.Transport{
			DisableKeepAlives: false, // Keep-alive is enabled
		},
	}

	// Configure the pusher
	p := &Pusher{
		Email:    "xdung24@gmail.com",
		Password: password,
		Host:     "https://api.openobserve.ai:443",
		Path:     "/api/dung_organization_20338_eul2VPBU0sHYNAe/default/_json",
		Client:   customClient,
	}

	// Start the logbeat loop
	go logbeatLoop(w, p, deviceName)

	// Wait for ctrl+C to exit
	log.Println("Press ctrl+C to exit...")
	c := make(chan os.Signal, 1)
	<-c
	log.Println("Exiting, good bye...")
}
