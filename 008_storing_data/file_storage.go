package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	data := []byte("Hello World")
	err := ioutil.WriteFile("file.txt", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("file.txt")
	fmt.Println(string(read1))

	file1, _ := os.Create("file2.txt")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Write %d bytes to file\n", bytes)

	file2, _ := os.Open("file2.txt")
	defer file2.Close()

	read2 := make([] byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d byte from file\n", bytes)
	fmt.Println(string(read2))

}
