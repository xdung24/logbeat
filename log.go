package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func readLastLine(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lastLine string

	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return lastLine, nil
}

func parseLog(text string) (timestamp string, port int, content string) {
	// Updated regular expression to match the format in the tests
	re := regexp.MustCompile(`\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})\](?: \[(\d+)\])? (.*)`)

	// Find matches
	matches := re.FindStringSubmatch(text)

	// Check if matches were found
	if len(matches) > 0 {
		timestamp = matches[1]            // The first group is always the timestamp
		content = matches[len(matches)-1] // The last group is always the content
		if len(matches) == 4 {            // If port is present
			num, err := strconv.Atoi(matches[2])
			if err != nil {
				num = 0 // Default port to 0 if conversion fails
			}
			port = num
		}
	}

	return
}
