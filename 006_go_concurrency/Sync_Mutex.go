package main

import (
	"sync"
	"fmt"
)

type AtomicInt struct {
	mu sync.Mutex // lock , only can locked by only one goroutine.
	n  int
}

/*
	Add method put n value to 'Atomic int'
 */
func (a *AtomicInt) Add(n int) {
	a.mu.Lock()
	a.n += n
	a.mu.Unlock()
}

func (a *AtomicInt) Value() int {
	a.mu.Lock()
	n := a.n
	a.mu.Unlock()
	return n
}

func main() {
	wait := make(chan struct{})
	var n AtomicInt
	go func() {
		n.Add(1) // one access
		close(wait)
	}()
	n.Add(1) // another concurrency acces
	<-wait
	fmt.Println(n.Value())
}
