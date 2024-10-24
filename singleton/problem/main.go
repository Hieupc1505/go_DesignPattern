package main

import (
	"log"
	"sync"
)

type config struct {
	logAllowed bool
}

func (c config) LogAllowed() bool {
	return c.logAllowed
}

func NewConfig(allowed bool) config {
	return config{logAllowed: allowed}
}

func main() {

	//Demo 1000 request to service at a same time
	//I made this code for simple demo, not a real practice

	rps := 1000
	wg := sync.WaitGroup{}
	wg.Add(rps)

	config := NewConfig(true)

	for i := 0; i < rps; i++ {
		go func(idx int) {
			requestHandler(idx)
			if config.LogAllowed() {
				log.Printf("Request %d handled successfully", idx)
			}
			wg.Done()

		}(i)
	}

}

func requestHandler(idx int) {
	//I have some log to print here
	// I have to know the config that if it was allowed or not
	// but I connot modify definition of this method
}
