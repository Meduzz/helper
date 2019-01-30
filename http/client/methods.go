package client

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type (
	HttpRequest struct {
		request *http.Request
	}

	HttpResponse struct {
		response *http.Response
	}
)

func GET(url string) (*HttpRequest, error) {
	return NewRequest("GET", url, nil)
}

func POST(url string, body interface{}) (*HttpRequest, error) {
	bs, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	req, err := NewRequest("POST", url, bytes.NewReader(bs))

	if err != nil {
		return nil, err
	}

	req.Header("Content-Type", "application/json")

	return req, nil
}

func PUT(url string, body interface{}) (*HttpRequest, error) {
	bs, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	req, err := NewRequest("PUT", url, bytes.NewReader(bs))

	if err != nil {
		return nil, err
	}

	req.Header("Content-Type", "application/json")

	return req, nil
}

func DELETE(url string, body interface{}) (*HttpRequest, error) {
	if body != nil {
		bs, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}
		return NewRequest("DELETE", url, bytes.NewReader(bs))
	} else {
		return NewRequest("DELETE", url, nil)
	}
}

func HEAD(url string) (*HttpRequest, error) {
	return NewRequest("HEAD", url, nil)
}

func NewRequest(method, url string, body io.Reader) (*HttpRequest, error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	return &HttpRequest{req}, nil
}

func (req *HttpRequest) Header(key, value string) {
	req.request.Header.Add(key, value)
}

func (req *HttpRequest) Do(client *http.Client) (*HttpResponse, error) {
	res, err := client.Do(req.request)

	if err != nil {
		return nil, err
	}

	return &HttpResponse{res}, nil
}

func (req *HttpRequest) Request() *http.Request {
	return req.request
}

func (res *HttpResponse) Code() int {
	return res.response.StatusCode
}

func (res *HttpResponse) Header(key string) string {
	return res.response.Header.Get(key)
}

func (res *HttpResponse) Body(target interface{}) error {
	defer res.response.Body.Close()
	bs, err := ioutil.ReadAll(res.response.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(bs, target)
}

func (res *HttpResponse) Response() *http.Response {
	return res.response
}
