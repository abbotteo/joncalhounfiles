package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://www.usegolang.com",
		"https://testwithgo.com",
		"https://invalid.gophercises.com",
	}
	var wg sync.WaitGroup
	responses := make([]string, len(urls))

	for i, url := range urls {
		wg.Add(1)
		// I am calling the index j so it is clear which variable we are using. You
		// could use variable shadowing here, or the "i := i" trick before even
		// writing the closure. Both would prevent someone from accidentally using
		// the wrong variable, but shadowing can introduce it's own set of
		// complications.
		go func(j int, url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				responses[j] = fmt.Sprintf("error: %v", err)
				return
			}
			responses[j] = resp.Status
		}(i, url)
	}
	wg.Wait()
}
