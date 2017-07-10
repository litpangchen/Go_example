package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	post := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	store(post, "file3.txt")
	var postRead Post
	load(&postRead, "file3.txt")
	fmt.Println(postRead)
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}
