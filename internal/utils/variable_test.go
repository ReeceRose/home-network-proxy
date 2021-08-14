package utils

import (
	"os"
	"testing"

	"github.com/ReeceRose/home-network-proxy/internal/consts"
	"github.com/stretchr/testify/assert"
)

// No OS wrapper used here as mocking requires too much boilerplate for this simple test
// and the changes of os.Setenv/os.Getenv/os.Clearenv not working are minimal

func TestVariable_GetVariable_ReturnsValueFromEnvironmentVariable(t *testing.T) {
	os.Clearenv()
	os.Setenv(consts.API_URL, "https://api.reecerose.com/api/v1")

	assert.Equal(t, "https://api.reecerose.com/api/v1", GetVariable(consts.API_URL))

	os.Clearenv()
}

func TestVariable_GetVariable_ReturnsDefaultValues(t *testing.T) {
	os.Clearenv()
	assert.Equal(t, "https://localhost:3000/api/v1/", GetVariable(consts.API_URL))
}
