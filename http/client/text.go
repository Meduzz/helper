package client

func POSTText(url string, body string, contentType string) (*HttpRequest, error) {
	return NewRequest("POST", url, []byte(body), contentType)
}

func PUTText(url string, body string, contentType string) (*HttpRequest, error) {
	return NewRequest("PUT", url, []byte(body), contentType)
}

func DELETEText(url string, body string, contentType string) (*HttpRequest, error) {
	return NewRequest("DELETE", url, []byte(body), contentType)
}
