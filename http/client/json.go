package client

import (
	"encoding/json"
)

func GET(url string) (*HttpRequest, error) {
	return NewRequest("GET", url, nil, EMPTY)
}

func POST(url string, body interface{}) (*HttpRequest, error) {
	bs, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	return NewRequest("POST", url, bs, JSON)
}

func PUT(url string, body interface{}) (*HttpRequest, error) {
	bs, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	return NewRequest("PUT", url, bs, JSON)
}

func DELETE(url string, body interface{}) (*HttpRequest, error) {
	if body != nil {
		bs, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}
		return NewRequest("DELETE", url, bs, JSON)
	} else {
		return NewRequest("DELETE", url, nil, EMPTY)
	}
}

func HEAD(url string) (*HttpRequest, error) {
	return NewRequest("HEAD", url, nil, EMPTY)
}
