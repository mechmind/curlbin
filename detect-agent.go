package main

import (
	"strings"
)

const (
	AGENT_CURL = iota
	AGENT_BROWSER
	AGENT_LIBRARY
	AGENT_OTHER
)

func detectAgent(userAgent string) int {
	switch {
	case strings.HasPrefix(userAgent, "curl"):
		return AGENT_CURL
	case strings.Contains(userAgent, "Mozilla"):
		return AGENT_BROWSER
	default:
		return AGENT_OTHER
	}
}
