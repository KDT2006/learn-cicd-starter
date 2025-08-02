package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIKey(t *testing.T) {
	value := "12344"
	validHeader := http.Header{
		"Authorization": []string{"ApiKey " + value},
	}
	result, err := GetAPIKey(validHeader)
	assert.NoError(t, err)
	assert.Equal(t, value, result)

	invalidHeader1 := http.Header{
		"Authorization": []string{"ApiKey"},
	}
	result, err = GetAPIKey(invalidHeader1)
	assert.Error(t, err)
	assert.Empty(t, result)

	invalidHeader2 := http.Header{}
	result, err = GetAPIKey(invalidHeader2)
	assert.Error(t, err)
	assert.Equal(t, ErrNoAuthHeaderIncluded, err)
	assert.Empty(t, result)
}
