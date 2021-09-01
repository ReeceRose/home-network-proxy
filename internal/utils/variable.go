package utils

import (
	"os"

	"github.com/ReeceRose/home-network-proxy/internal/consts"
)

// GetVariable returns a value given a key. It will first try to read from environment variables and will default to preset values
func GetVariable(key string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return getDefaultForKey(key)
}

// getDefaultForKey is a handy method to get the default values if not present in environment variables
func getDefaultForKey(key string) string {
	switch key {
	case consts.API_URL:
		return "https://p7xsh0rld4.execute-api.us-east-1.amazonaws.com/production"
	}
	return ""
}
