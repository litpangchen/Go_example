package main

//import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello!"
		close(ch)
	}()

	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	//v, ok := <-ch
	//fmt.Println(v)
	//fmt.Println(ok)
}
