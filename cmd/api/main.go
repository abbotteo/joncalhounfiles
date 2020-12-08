package main

import (
	"fmt"
	"math/rand"
	"time"
)

func notify(services ...string) {
	res := make(chan string)
	count := 0

	for _, service := range services {
		count++
		go func(s string) {
			fmt.Printf("Starting to notifing %s...\n", s)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			res <- fmt.Sprintf("Finished %s", s)
		}(service)
	}

	for i := 0; i < count; i++ {
		fmt.Println(<-res)
	}

	fmt.Println("All services notified!")
}

func main() {
	notify("Service-1", "Service-2", "Service-3")
	// Running this outputs "All services notified!" but we
	// won't see any of the services outputting their finished messages!
}