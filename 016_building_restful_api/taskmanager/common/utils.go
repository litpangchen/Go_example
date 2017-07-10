package common

import (
	"os"
	"log"
	"encoding/json"
	"net/http"
)

type AppError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	HttpStatus int `json:"status"`
}

type ErrorResource struct {
	Data AppError `json:"data"`
}

type Configuration struct {
	Server, MongoDbHost, DBUser, DBPwd, Database string
}

var AppConfig Configuration

func initConfig() {
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("common/config.json")

	defer file.Close()

	if err != nil {
		log.Fatalf("load config : %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = Configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("load app config : %s\n", err)
	}
}

func DisplayAppError(w http.ResponseWriter, handleError error, message string, code int) {
	errObj := AppError{
		Error:      handleError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if json, err := json.Marshal(ErrorResource{Data: errObj}); err == nil {
		w.Write(json)
	}
}
