package main

import "fmt"

type Sushi string

func main() {

	var ch <-chan Sushi = Producer()
	for s := range ch {
		fmt.Println("Consumed ", s)
	}
}

func Producer() <-chan Sushi {
	ch := make(chan Sushi)

	go func() {
		ch <- Sushi("Sushi A")
		ch <- Sushi("Sushi B")
		close(ch)
	}()

	go func() {
		ch <- Sushi("Sushi C")
		ch <- Sushi("Sushi D")
		close(ch)
	}()

	return ch
}
