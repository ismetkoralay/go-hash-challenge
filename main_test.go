package main

import (
	"go-hash-challenge/model"
	"testing"
)

var dummy = [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6}

type testHashService struct {
}

func (testHashService) CalculateHashWorker(url string, ch chan model.UrlResponse) {
	ch <- model.UrlResponse{Url: url, ResponseHash: dummy}
}

func TestGetResponseHashesSingleInputCount(t *testing.T) {
	result, _ := GetResponseHashes(testHashService{}, 1, []string{"google.com"})
	if len(result) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result))
	}
}

func TestGetResponseHashesNoArgs(t *testing.T) {
	_, err := GetResponseHashes(testHashService{}, 1, []string{})
	if err == nil {
		t.Error("Tried to run with no args")
	}
}

func TestGetResponseHashesSingleInputValue(t *testing.T) {
	result, _ := GetResponseHashes(testHashService{}, 1, []string{"google.com"})
	if result["google.com"] != dummy {
		t.Errorf("Expected %v, got %v", dummy, result["http://google.com"])
	}
}

func TestGetResponseHashesMultipleInputsCount(t *testing.T) {
	args := []string{"http://google.com", "http://adjust.com"}
	result, _ := GetResponseHashes(testHashService{}, 1, args)
	if len(result) != len(args) {
		t.Errorf("Expected %d result, got %d", len(args), len(result))
	}
}

func TestGetResponseHashesMultipleInputsValue(t *testing.T) {
	args := []string{"http://google.com", "http://adjust.com"}
	result, _ := GetResponseHashes(testHashService{}, 1, args)
	for _, val := range args {
		if result[val] != dummy {
			t.Errorf("For %s expected %v got %v", val, dummy, result[val])
		}
	}
}

func TestGetResponseHashesZeroParallel(t *testing.T) {
	args := []string{"http://google.com"}
	_, err := GetResponseHashes(testHashService{}, 0, args)
	if err == nil {
		t.Errorf("0 Parallel arg expected error but didn't get")
	}
}
