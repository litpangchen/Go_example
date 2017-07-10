package main

import "fmt"

func main() {
	fmt.Println(calculate(2, 3))
}

func calculate(x, y int) (total int) {
	return x + y;
}
