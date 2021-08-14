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
	case consts.API_PORT:
		return "3000"
	case consts.API_URL:
		return "https://localhost:3000/api/v1/"
	case consts.CERT_DIR:
		return "certs"
	case consts.CLIENT_CERT:
		return "localhost.crt"
	case consts.API_CERT:
		return "localhost.crt"
	case consts.API_KEY:
		return "localhost.key"
	case consts.DB_URI:
		return "mongodb://localhost:27017/" + GetVariable(consts.DB_NAME)
	case consts.DB_NAME:
		return "home-network-proxy"
	case consts.DB_USER:
		return "admin"
	case consts.DB_PASS:
		return "admin"
	}
	return ""
}
