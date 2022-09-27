package main

import "fmt"

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
Usage

To help diagnose such bugs, Go includes a built-in data race detector. To use it, add the -race flag to the go command:

$ go test -race mypkg    // to test the package
$ go run -race mysrc.go  // to run the source file
*/
