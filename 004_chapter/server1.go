package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"time"
	"encoding/base64"
)

type Post struct {
	User    string
	Threads []string
}

func main() {
	server := http.Server{
		Addr:"127.0.0.1:8001",
	}
	http.HandleFunc("/index", index)
	http.HandleFunc("/header", header)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/json", jsonExample)
	http.HandleFunc("/processForm", processForm)
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	http.HandleFunc("/showMessage", showMessage)
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func header(w http.ResponseWriter, r*http.Request) {
	h := r.Header["Content-Type"]
	fmt.Fprintln(w, h)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r*http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User : "lucas chen",
		Threads:[]string{
			"first",
			"second",
			"third",
		},

	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func processForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:"first_cookie",
		Value:"Go Web Programming",
		HttpOnly:true,
	}

	c2 := http.Cookie{
		Name:"second_cookie",
		Value:"Go Web Programming 2",
		HttpOnly:true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r*http.Request) {
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the furst cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func showMessage(w http.ResponseWriter, r*http.Request) {
	c, err := r.Cookie("first_cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No cookie found")
		}
	} else {
		rc := http.Cookie{
			Name:"Flash",
			MaxAge: -1,
			Expires:time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}
