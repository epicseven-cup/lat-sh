package latsh

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/epicseven-cup/lat-sh/internal/request"
	"io"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

type LatShClient struct {
	UserAgent          string
	MaxIdleConns       int
	IdleConnTimeout    time.Duration
	DisableCompression bool
	_client            *http.Client
	Headers            map[string]string
}

func NewDefaultLatShClient() *LatShClient {
	client := http.DefaultClient
	return &LatShClient{
		_client: client,
	}
}

func NewLatShClient(userAgent string, header map[string]string, timeout time.Duration, maxIdleConns int, idleConnTimeout time.Duration, disableCompression bool) *LatShClient {
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

	post, err := c._client.Post(url, contentType, io.Reader(bytes.NewBuffer(jsonData)))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(post.Body)
	if post.StatusCode != http.StatusCreated {
		return errors.New(post.Status)
	}

	jsonResponse := Response{}
	err = json.NewDecoder(post.Body).Decode(&jsonResponse)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Response status code: %d", post.StatusCode))
	fmt.Println(fmt.Sprintf("Message: %s", jsonResponse.Message))
	return nil
}
