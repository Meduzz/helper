package client

import (
	"strings"
)

func POSTText(url string, body string, contentType string) (*HttpRequest, error) {
	return NewRequest("POST", url, strings.NewReader(body), contentType)
}

func PUTText(url string, body string, contentType string) (*HttpRequest, error) {
	return NewRequest("PUT", url, strings.NewReader(body), contentType)
}

func DELETEText(url string, body string, contentType string) (*HttpRequest, error) {
	return NewRequest("DELETE", url, strings.NewReader(body), contentType)
}
