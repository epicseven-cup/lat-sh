package client

import (
	"encoding/json"
	"github.com/epicseven-cup/lat-sh/internal/request"
	"io"
	"net/http"
	"time"
)

type LatShClient struct {
	UserAgent          string
	MaxIdleConns       int
	IdleConnTimeout    time.Duration
	DisableCompression bool
	_client            *http.Client
	Headers            map[string]string
}

func NewClient(userAgent string, header map[string]string, timeout time.Duration, maxIdleConns int, idleConnTimeout time.Duration, disableCompression bool) *LatShClient {
	client := http.DefaultClient
	client.Timeout = timeout
	client.Transport = &http.Transport{
		DisableCompression: disableCompression,
		MaxIdleConns:       maxIdleConns,
		IdleConnTimeout:    idleConnTimeout,
	}
	return &LatShClient{
		UserAgent: userAgent,
		Headers:   header,
		_client:   client,
	}
}

func (c *LatShClient) CreateUrl(source string, destination string) error {
	url, contentType, createUrlRequest := request.NewCreateUrl(source, destination)
	jsonData, err := json.Marshal(createUrlRequest)
	if err != nil {
		return err
	}
	
	c._client.Post(url, contentType, io.Reader()
}
