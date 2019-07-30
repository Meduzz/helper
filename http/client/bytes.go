package client

func POSTBytes(url string, body []byte, contentType string) (*HttpRequest, error) {
	return NewRequest("POST", url, body, contentType)
}

func PUTBytes(url string, body []byte, contentType string) (*HttpRequest, error) {
	return NewRequest("PUT", url, body, contentType)
}

func DELETEBytes(url string, body []byte, contentType string) (*HttpRequest, error) {
	return NewRequest("DELETE", url, body, contentType)
}
