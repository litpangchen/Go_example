package common

import (
	"net/http"
	"crypto/rsa"
	"io/ioutil"
	"log"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
	"strings"
	"github.com/go-errors/errors"
)

type AppClaims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func initKeys() {
	var err error
	signKeyBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("Read private Key file : %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signKeyBytes)
	if err != nil {
		log.Fatalf("Parse private key : %s\n", err)
	}
	verifyKeyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("Read public Key file : %s\n", err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyKeyBytes)
	if err != nil {
		log.Fatalf("Parse public key : %s\n", err)
	}
}

func GenerateJWT(name, role string) (string, error) {
	claims := &AppClaims{
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequestWithClaims(
		r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				DisplayAppError(w, err, "Access token is expired, get a new token", 401)
				return
			default:
				DisplayAppError(w, err, "Error parsing the Access Token", 500)
				return
			}
		default:
			DisplayAppError(w, err, "Error parsing the Access Token", 500)
			return
		}
	}

	if token.Valid {
		context.Set(r, "user", token.Claims.(*AppClaims).UserName)
		next(w, r)
	} else {
		DisplayAppError(w, err, "Invalid Access Token", 401)
	}
}

func TokenFromAuthHeader(r *http.Request) (string, error) {
	if authorizationHeader := r.Header.Get("Authorization"); authorizationHeader != "" {
		if len(authorizationHeader) > 6 && strings.ToUpper(authorizationHeader[0:6]) == "BEARER" {
			return authorizationHeader[:7], nil
		}
	}
	return "", errors.New("No token in the HTTP request")
}
