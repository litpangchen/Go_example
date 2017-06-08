package main

import (
	"sync"
	"fmt"
)

func main() {
	people := []string{"Ana", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1)
	wg := new(sync.WaitGroup)
	wg.Add(len(people))
	for _, name := range people {
		go Seek(name, match, wg)
	}
	wg.Wait()
	select {
	case name := <-match:
		{
			fmt.Printf("No one received %s 's message\n", name)
		}
	default:
	}
}

func Seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		{
			fmt.Printf("%s sent a message to %s.\n", peer, name)
		}
	case match <- name:
		{

		}
	}
	wg.Done()
}
