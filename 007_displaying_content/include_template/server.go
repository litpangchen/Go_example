package main

import (
	"net/http"
	"html/template"
	"time"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	http.HandleFunc("/date", processDate)

	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("t1.html", "t2.html"))
	t.Execute(w, nil)
}

func formatDate(t time.Time) string {
	return t.Format("2006-01-02 23:59:59")
}

func processDate(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fDate": formatDate}
	t := template.New("date.html").Funcs(funcMap)
	t, _ = t.ParseFiles("date.html")
	t.Execute(w, time.Now())

}
