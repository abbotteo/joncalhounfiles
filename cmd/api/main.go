package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func notify(wg *sync.WaitGroup, services ...string) {
	res := make(chan string, len(services))
	wg.Add(len(services))
	count := 0
	for _, service := range services {
		count++
		go func(s string) {
			fmt.Printf("Starting to notifing %s...\n", s)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			res <- fmt.Sprintf("Finished %s", s)
			wg.Done()
		}(service)
	}

	response := <-res
	fmt.Println(response)
	fmt.Println("One service notified!")
}

func main() {
	var wg sync.WaitGroup
	notify(&wg, "Service-1", "Service-2", "Service-3", "Service-1b", "Service-2b", "Service-3b", "Service-1c", "Service-2c", "Service-3c")
	// Running this outputs "All services notified!" but we
	// won't see any of the services outputting their finished messages!
	fmt.Printf("routines %v...\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("routines %v...\n", runtime.NumGoroutine())
}
