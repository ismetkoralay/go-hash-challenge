package service

import (
	"crypto/md5"
	"errors"
	"go-hash-challenge/model"
	"testing"
)

var dummy = []byte("")

type testClient struct {
}

func (testClient) Get(url string) ([]byte, error) {
	if url == "http://google.com" {
		return dummy, nil
	} else {
		return nil, errors.New("No body")
	}
}

func TestHashService_CalculateHashWorker(t *testing.T) {
	hashService := HashService{Client: testClient{}}
	ch := make(chan model.UrlResponse)
	go hashService.CalculateHashWorker("http://google.com", ch)

	res := <-ch

	expected := md5.Sum(dummy)
	if res.ResponseHash != expected {
		t.Errorf("Expected %v but got %v", expected, res.ResponseHash)
	}

	if res.Url != "http://google.com" {
		t.Errorf("Expected http://google.com but got %s", res.Url)
	}
}

func TestHashService_CalculateHashWorkerDummyUrl(t *testing.T) {
	hashService := HashService{Client: testClient{}}
	ch := make(chan model.UrlResponse)
	go hashService.CalculateHashWorker("http://dummy.com", ch)

	res := <-ch
	var expected [16]byte

	if res.ResponseHash != expected {
		t.Errorf("zxczcz %v", res.ResponseHash)
	}
}
