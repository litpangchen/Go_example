package main

import (
	"fmt"
	"time"
)

func main() {
	//go fmt.Println("Hello another go")
	//fmt.Println("Hello")
	//time.Sleep(time.Second)

	Publish("go is the best !", time.Second * 5)
	fmt.Println("Let hope the news will be publish before I leave!")
	time.Sleep(10 * time.Second)
	fmt.Println("Ten second later i am leave !")

}

func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("Breaking news ! ", text)
	}()
}
