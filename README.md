# go-hash-challenge

It's a simple project that send request to urls that given by the command line arguments(./go-hash-challenge google.com github.com) and prints the hashes of the response bodies. It also takes the max number of parallel processes like that ./go-hash-challenge -parallel 5 google.com github.com(the default is 10).

To able to run the project you might follow these instructions:

* First you should have go. (https://golang.org/doc/install)
* Then pull there repository. (git pull https://github.com/ismetkoralay/go-hash-challenge.git)
* Then run the commands at the root of the project.
  - go mod tidy
  - go build
* To run the test go test ./... (It runs all the test including the ones in the service/hashservice_test.go)
* Then at the same folder just run it like that ./go-hash-challenge -parallel 5 google.com github.com
