package main

import (
	"net/http"
	"math/rand"
	"html/template"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()

}

func process(w http.ResponseWriter, r *http.Request) {
	randNum := rand.Intn(10)

	var t *template.Template

	if randNum > 5 {
		t, _ = template.ParseFiles("layout.html", "blue_hello.html")
	} else {
		t, _ = template.ParseFiles("layout.html", "yellow_hello.html")
	}
	t.ExecuteTemplate(w, "layout", nil)
}
