package main

import "fmt"

func main() {

	/*
		Goroutine between goroutine sharing data through channel.
	 */

	ch := make(chan int)
	go func() {
		n := 0
		n++
		ch <- n
	}()

	n := <-ch
	n++
	fmt.Println(n)
}
