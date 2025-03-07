package rpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestEncoding struct {
	Testing bool `json:"testing"`
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"testing\":true}"
	actual := EncodeMessage(TestEncoding{Testing: true})
	assert.Equal(t, expected, actual)
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"method\":\"hi\"}"
	method, content, err := DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)
	assert.NoError(t, err)
	assert.Equal(t, 15, contentLength)
	assert.Equal(t, "hi", method)
}
