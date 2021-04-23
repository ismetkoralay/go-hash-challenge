package service

import (
	"crypto/md5"
	"go-hash-challenge/client"
	"go-hash-challenge/model"
	"strings"
)

type IHashService interface {
	CalculateHashWorker(url string, ch chan model.UrlResponse)
}

type HashService struct {
	Client client.IHttpClient
}

func (hS HashService) CalculateHashWorker(url string, ch chan model.UrlResponse) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	body, err := hS.Client.Get(url)
	if err != nil {
		ch <- model.UrlResponse{
			Url: url,
		}
	} else {
		hashByte := md5.Sum(body)
		ch <- model.UrlResponse{
			Url:          url,
			ResponseHash: hashByte,
		}
	}
}
