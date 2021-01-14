package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://invalid.gophercises.com",
		"https://www.usegolang.com",
		"https://testwithgo.com",
		"https://www.abbotteo.com.ar",
	}
	var wg sync.WaitGroup
	responses := make([]string, len(urls))
	racingOrder := 0

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
			racingOrder++
			if err != nil {
				responses[j] = fmt.Sprintf("error: %v - order: %v", err, racingOrder)
				return
			}
			responses[j] = fmt.Sprintf("%v - order: %v", resp.Status, racingOrder)
		}(i, url)
	}
	wg.Wait()

	for i, url := range urls {
		fmt.Printf("%v: %v\n", url, responses[i])
	}
}
