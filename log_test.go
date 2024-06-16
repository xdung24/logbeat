package main

import (
	"testing"
)

type test struct {
	text    string
	time    string
	port    int
	content string
}

func Test_ParseLog(t *testing.T) {
	tests := []test{
		{
			text:    "[2024-06-16 16:39:46] [1234] Found 1 boss rallies",
			time:    "2024-06-16 16:39:46",
			port:    1234,
			content: "Found 1 boss rallies",
		},
		{
			text:    "[2024-06-16 16:39:46] [65843] Found 2 boss rallies",
			time:    "2024-06-16 16:39:46",
			port:    65843,
			content: "Found 2 boss rallies",
		},
		{
			text:    "[2024-06-16 16:39:46] Found 1 boss rallies",
			time:    "2024-06-16 16:39:46",
			port:    0,
			content: "Found 1 boss rallies",
		},
	}

	for _, test := range tests {
		time, port, content := parseLog(test.text)
		t.Logf("Time: %s, Port: %v, Content: %s\n", time, port, content)
		t.Logf("Expected Time: %s, Port: %v, Content: %s\n", test.time, test.port, test.content)
		if time != test.time || port != test.port || content != test.content {
			t.Errorf("Failed to parse log: %s\n", test.text)
		}
	}
}
