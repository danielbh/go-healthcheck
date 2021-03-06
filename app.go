// Go contains rich function for grab web contents. _net/http_ is the major
// library.
// Ref: [golang.org](http://golang.org/pkg/net/http/#pkg-examples).
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	var url = "https://golangcode.com"

	client := http.Client{
		Timeout: time.Duration(5000 * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// Print the HTTP Status Code and Status Name
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("HTTP Status is in the 2xx range")
	} else {
		fmt.Println("Argh! Broken")
	}
}
