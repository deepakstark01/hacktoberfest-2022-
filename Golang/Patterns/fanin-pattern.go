package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Multiplexed two channels into one
	multiplexer := getMultiplexer(getDataChannel(1), getDataChannel(3))
	// Multiplexed two channels into one
	ch := getMultiplexer(multiplexer, getDataChannel(2))
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}

func getMultiplexer(input ...<-chan interface{}) <-chan interface{} {
	ch := make(chan interface{})
	for i := 0; i < len(input); i++ {
		inCh := input[i]
		go func() {
			for {
				ch <- <-inCh
			}
		}()
	}
	return ch
}

// getDataChannel is generator function
func getDataChannel(id int) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for true {
			ch <- fmt.Sprintf("%d got from %d", rand.Intn(200), id)
		}
	}()
	return ch
}