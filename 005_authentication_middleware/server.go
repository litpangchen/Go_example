package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"crypto"
)

type Book struct {
	Id    int
	Name  string
	Price float32
}

func main() {
	server := http.Server{
		Addr:"0.0.0.0:8000",
	}
	bookHandler := http.HandlerFunc(book)
	http.Handle("/book", auth2(bookHandler))

	server.ListenAndServe()
}

func auth2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signer := jwt.SigningMethodHMAC{
			Name:"lol",
			Hash:crypto.SHA256,
		}
		signature, _ := signer.Sign("text", []byte("password123"))
		fmt.Println(signature)
		err := signer.Verify("text", signature, []byte("password123"))

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func book(w http.ResponseWriter, r *http.Request) {

	var books = [] Book{
		Book{Id:1, Name:"book1", Price:10.0},
		Book{Id:2, Name:"book2", Price:20.0},
	}
	w.Header().Set("Content-Type", "application/json")
	byte, _ := json.Marshal(books)
	w.Write(byte)
}


