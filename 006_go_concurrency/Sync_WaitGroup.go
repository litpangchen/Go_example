package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {

		n := i
		go func() {
			fmt.Println(n)
			wg.Done()
		}()

		/*
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
		*/
	}
	wg.Wait()
	fmt.Println()
}
