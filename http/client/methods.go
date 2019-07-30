package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type (
	HttpRequest struct {
		request *http.Request
		body    []byte
	}

	HttpResponse struct {
		response *http.Response
	}
)

const (
	JSON  = "application/json"
	TEXT  = "text/plain"
	HTML  = "text/html"
	BYTES = "application/octet-stream"
	EMPTY = ""
)

func NewRequest(method, url string, body []byte, contentType string) (*HttpRequest, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	wrap := &HttpRequest{req, body}
	if contentType != "" {
		wrap.Header("Content-Type", contentType)
	}

	return wrap, nil
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

func (req *HttpRequest) Sign(signer func(req *HttpRequest) error) error {
	return signer(req)
}

func (req *HttpRequest) Body() []byte {
	return req.body
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
