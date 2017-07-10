package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"crypto/rsa"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type Response struct {
	Text string `json:"text"`
}

const (
	privateKeyPath = "app.rsa"
	publicKeyPath  = "app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func init() {
	var err error

	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	token, err := request.ParseFromRequest(
		r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Token expired,get a new one")
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Error while parsing token")
				return
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while parsing token")
			return
		}
	}

	if token.Valid {
		response := Response{"Authorized to System"}
		jsonResponse(response, w)
	} else {
		response := Response{"Invalid token"}
		jsonResponse(response, w)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error in request body")
		return
	}
	if user.Username != "lucas" && user.Password != "password123" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Wrong info")
		return
	}

	t := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["CustomerUserInfo"] = struct {
		Name string
		Role string
	}{user.Username, "Member"}
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	t.Claims = claims
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Token signing error:%v\n", err)
		return
	}

	response := Token{tokenString}
	jsonResponse(response, w)

}

func jsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/auth", authHandler).Methods("POST")
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Server started on port 8080")
	server.ListenAndServe()
}
