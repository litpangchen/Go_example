package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"fmt"
	"os"
)

/*
go test -run ''      # Run all tests.
go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
 */

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	fmt.Println("Code : ", code)
	tearDown()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/posts/", handleRequest)
	writer = httptest.NewRecorder()
}

func tearDown() {
	fmt.Println("Tear Down")
}

func TestHandleGet(t *testing.T) {

	request, _ := http.NewRequest("GET", "/posts/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	fmt.Println(writer.Body)
	if post.Id != 1 {
		t.Error("Cannot retrieve Json POST")
	}
}

func TestHandlePut(t *testing.T) {

}
