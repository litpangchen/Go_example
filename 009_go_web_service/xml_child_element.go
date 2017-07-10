package main

import (
	"encoding/xml"
	"os"
	"io/ioutil"
	"fmt"
)

type Post struct {
	XMLName  xml.Name `xml:"post"`
	Id       string `xml:",attr"`
	Content  string `xml:"content"`
	Author   Author `xml:"author"`
	Xml      string `xml:",innerxml""`
	Comments [] Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:author`
}

func main() {
	xmlFile, err := os.Open("xml2.xml")
	if err != nil {
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)

}
