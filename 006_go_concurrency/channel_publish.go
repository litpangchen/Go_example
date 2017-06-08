package main

import (
	"time"
	"fmt"
)

func main() {
	wait := Publish("Channnels let goroutines communicate", 5 * time.Second)
	<-wait
	fmt.Println("The news is out, time to leave")
}

func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS :", text)
		//close(ch)
	}()
	return ch
}
