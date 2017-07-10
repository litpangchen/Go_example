package main

import (
	"time"
	"fmt"
)

/*
Unbuffered channels are synchronous.
You can think of unbuffered channels as a box that can contain only one thing at a time.
Once a goroutine puts something into this box, no other goroutines can put anything in,
unless another goroutine takes out whatever is inside it first.
This means if another goroutine wants to put in something else when the box contains something already,
it will block and go to sleep until the box is empty.
Similarly, if a goroutine tries to take out something from this box and it’s empty,
it’ll block and go to sleep until the box has something in it.
The syntax for putting things into a channel is quickly recognizable, visually.
This puts an integer 1 into the channel ch:
ch <- 1

Taking out the value from a channel is equally recognizable.
This removes the value from the channel and assigns it to the variable i:
i := <- ch

Channels can be directional. By default, channels work both ways (bidirectional)
and values can be sent to or received from it. But channels can be restricted to send-only or receive-only.

This allocates a send-only channel of strings:
        ch := make(chan <- string)

This allocates a receive-only channel of strings:
        ch := make(<-chan string)

Buffered channels are asynchronous, first-in, first-out (FIFO) message queues.
Think of buffered channels as a large box that can contain a number of similar things.
A goroutine can continually add things into this box without blocking until there’s no more space in the box.
Similarly, another goroutine can continually remove things from this box (in the same sequence it was put in)
and will only block when it runs out of things to remove.
 */

func printNumber2(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	w <- true
}

func printLetter2(w chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	w <- true
}

func main() {
	w1, w2 := make(chan bool), make(chan bool)
	go printNumber2(w1)
	go printLetter2(w2)
	fmt.Println(<-w1)
	fmt.Println(<-w2)

}
