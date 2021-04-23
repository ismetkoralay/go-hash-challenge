package main

import (
	"errors"
	"flag"
	"fmt"
	"go-hash-challenge/client"
	"go-hash-challenge/model"
	"go-hash-challenge/service"
	"sync"
)

func main() {
	parallel := flag.Int("parallel", 10, "max parallel requests allowed")
	flag.Parse()
	hashService := service.HashService{
		Client: client.HttpClient{},
	}
	_, err := GetResponseHashes(hashService, *parallel, flag.Args())
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetResponseHashes(hashService service.IHashService, parallel int, args []string) (map[string][16]byte, error) {
	if len(args) < 1 {
		return nil, errors.New("No url provided")
	}

	if parallel <= 0 {
		return nil, errors.New("Can't run with 0 concurrent threads.")
	}

	urls := make(chan string)

	go func() {
		defer close(urls)
		for _, u := range args {
			urls <- u
		}
	}()

	var wg sync.WaitGroup

	results := make(chan model.UrlResponse)

	for i := 0; i < parallel; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range urls {
				hashService.CalculateHashWorker(url, results)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var response = make(map[string][16]byte)

	for res := range results {
		var hash [16]byte
		if len(res.ResponseHash) > 0 {
			hash = res.ResponseHash
			fmt.Printf("%s %x \n", res.Url, res.ResponseHash)
		} else {
			fmt.Printf("%s No Response \n", res.Url)
		}
		response[res.Url] = hash
	}
	return response, nil
}
