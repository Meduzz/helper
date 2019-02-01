package client

import (
	"bytes"
)

func POSTBytes(url string, body []byte, contentType string) (*HttpRequest, error) {
	return NewRequest("POST", url, bytes.NewReader(body), contentType)
}

func PUTBytes(url string, body []byte, contentType string) (*HttpRequest, error) {
	return NewRequest("PUT", url, bytes.NewReader(body), contentType)
}

func DELETEBytes(url string, body []byte, contentType string) (*HttpRequest, error) {
	return NewRequest("DELETE", url, bytes.NewReader(body), contentType)
}
