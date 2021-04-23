package client

import (
	"io/ioutil"
	"net/http"
)

type IHttpClient interface {
	Get(url string) ([]byte, error)
}

type HttpClient struct {
}

func (client HttpClient) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
