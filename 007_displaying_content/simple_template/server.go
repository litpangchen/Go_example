package main

import (
	"net/http"
	"html/template"
	"math/rand"
	"time"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/range", processRange)
	http.HandleFunc("/with", with)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("tmpl.html"))
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(10)
	//t, _ := template.Must(template.ParseGlob("*html"))
	//t.ExecuteTemplate(w, "tmpl.html", "Hello World")
	t.Execute(w, randNum > 5)
}

func processRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("range.html"))
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func with(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("with.html"))
	t.Execute(w, "hello")
}
