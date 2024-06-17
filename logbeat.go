package main

import (
	"log"
	"time"

	"github.com/radovskyb/watcher"
)

// Map to store the last line per file
var lastlines = make(map[string]string)

func logbeatLoop(w *watcher.Watcher, p *Pusher, deviceName string) {

	go func() {
		for {
			select {
			case event := <-w.Event:
				// Read the last line
				lastLine, err := readLastLine(event.Path)
				if err != nil {
					log.Println(err)
					continue
				}

				// Parse the last line
				// Push the log to the server
				go func() {
					timestamp, port, content := parseLog(lastLine)
					log.Printf("Pushing: %s %s %s %d %s\n", event.Name(), timestamp, deviceName, port, content)
					p.pushLog(event.Name(), timestamp, content, deviceName, port)
				}()

				// If the last line has changed
				if lastLine != lastlines[event.Path] {
					// Update the line count map
					lastlines[event.Path] = lastLine
				}
			case err := <-w.Error:
				log.Println(err)
			case <-w.Closed:
				log.Println("Watcher closed")
			}
		}
	}()

	// Print a list of all files and folders currently being watched.
	for path, f := range w.WatchedFiles() {
		log.Printf("Watching: %s (IsDir: %v)\n", path, f.IsDir())
	}

	// Start watching.
	if err := w.Start(time.Millisecond * 1000); err != nil {
		log.Fatalln(err)
	}
}
