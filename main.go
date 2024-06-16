package main

import (
	"log"
	"os"

	"github.com/radovskyb/watcher"
)

func main() {
	// Get the device name
	var deviceName = getDevice()

	// Create a new file watcher.
	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Write)

	// Watch a specific folder recursively for changes.
	folder := `C:\Users\dung\EvonyBot\logs\`
	if err := w.AddRecursive(folder); err != nil {
		log.Fatalln(err)
	}

	// Configure the pusher
	p := &Pusher{
		Email:    "xdung24@gmail.com",
		Password: "1hn3O08WS7ZR65MU4d29",
		Host:     "https://api.openobserve.ai:443",
		Path:     "/api/dung_organization_20338_eul2VPBU0sHYNAe/default/_json",
		Timeout:  15,
	}

	// Start the logbeat loop
	go logbeatLoop(w, p, deviceName)

	// Wait for ctrl+C to exit
	log.Println("Press ctrl+C to exit...")
	c := make(chan os.Signal, 1)
	<-c
	log.Println("Exiting, good bye...")
}
