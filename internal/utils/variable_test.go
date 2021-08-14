package utils

import (
	"os"
	"testing"

	"github.com/ReeceRose/home-network-proxy/internal/consts"
	"github.com/stretchr/testify/assert"
)

func TestVariable_GetVariable_ReturnsValueFromEnvironmentVariable(t *testing.T) {
	os.Clearenv()
	os.Setenv(consts.API_PORT, "9000")
	os.Setenv(consts.API_URL, "https://api.reecerose.com/api/v1")
	os.Setenv(consts.CERT_DIR, "/var/certs/")
	os.Setenv(consts.API_CERT, "api.crt")
	os.Setenv(consts.API_KEY, "api.key")
	os.Setenv(consts.DB_URI, "mongodb://localhost:27017/hnp")
	os.Setenv(consts.DB_NAME, "hnp")
	os.Setenv(consts.DB_USER, "root")
	os.Setenv(consts.DB_PASS, "pass123")

	assert.Equal(t, "9000", GetVariable(consts.API_PORT))
	assert.Equal(t, "https://api.reecerose.com/api/v1", GetVariable(consts.API_URL))
	assert.Equal(t, "/var/certs/", GetVariable(consts.CERT_DIR))
	assert.Equal(t, "api.crt", GetVariable(consts.API_CERT))
	assert.Equal(t, "api.key", GetVariable(consts.API_KEY))
	assert.Equal(t, "mongodb://localhost:27017/hnp", GetVariable(consts.DB_URI))
	assert.Equal(t, "hnp", GetVariable(consts.DB_NAME))
	assert.Equal(t, "root", GetVariable(consts.DB_USER))
	assert.Equal(t, "pass123", GetVariable(consts.DB_PASS))
	assert.Equal(t, "", GetVariable("unknown"))

	os.Clearenv()
}

func TestVariable_GetVariable_ReturnsDefaultValues(t *testing.T) {
	os.Clearenv()
	assert.Equal(t, "3000", GetVariable(consts.API_PORT))
	assert.Equal(t, "https://localhost:3000/api/v1/", GetVariable(consts.API_URL))
	assert.Equal(t, "certs", GetVariable(consts.CERT_DIR))
	assert.Equal(t, "localhost.crt", GetVariable(consts.API_CERT))
	assert.Equal(t, "localhost.key", GetVariable(consts.API_KEY))
	assert.Equal(t, "mongodb://localhost:27017/home-network-proxy", GetVariable(consts.DB_URI))
	assert.Equal(t, "home-network-proxy", GetVariable(consts.DB_NAME))
	assert.Equal(t, "admin", GetVariable(consts.DB_USER))
	assert.Equal(t, "admin", GetVariable(consts.DB_PASS))
	assert.Equal(t, "", GetVariable("unknown"))
}
