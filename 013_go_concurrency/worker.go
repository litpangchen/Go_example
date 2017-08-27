package main

import (
	"fmt"
	"sync"
	"time"
	"strconv"
)

var waitGroup sync.WaitGroup
var data chan string

func main() {
	fmt.Println("Starting the application")
	data = make(chan string)

	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go worker(i)
	}

	for i := 0; i < 50; i++ {
		data <- "Testing " + strconv.Itoa(i)
	}

	close(data)
	waitGroup.Wait()
}

func worker(workerNum int) {
	fmt.Println("Goroutine working is now starting")
	defer func() {
		fmt.Println("Destroying worker")
		waitGroup.Done()
	}()

	for {
		value, ok := <-data
		if !ok {
			fmt.Println("The channel is closed")
			break
		}
		fmt.Println("Worker " + strconv.Itoa(workerNum) + " received " + value)
		time.Sleep(time.Second * 1)
	}
}
