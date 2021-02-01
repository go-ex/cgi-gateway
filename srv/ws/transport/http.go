package transport

import (
	"bytes"
	"net/http"
)

type httpTransport struct {
}

func (h *httpTransport) NewRequest(url string) Request {
	return &httpClient{
		url: url,
	}
}

func NewHttpTransport() *httpTransport {
	return &httpTransport{}
}

type httpClient struct {
	url string
}

func (request *httpClient) Call(message []byte) error {
	resp, err := http.Post(
		request.url,
		"application/json",
		bytes.NewReader(message),
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
