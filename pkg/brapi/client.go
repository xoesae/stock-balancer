package brapi

import (
	"net/http"
	"time"
)

func newHttpClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}

func NewBrapiClient(token string, url string) Brapi {
	return Brapi{
		Token:  token,
		Url:    url,
		client: newHttpClient(10 * time.Second),
	}
}
