package main

import (
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

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				// Make note of an error
				return
			}
			// Make note of the resp.Status
			_ = resp.Status
		}(url)
	}
	wg.Wait()
}
