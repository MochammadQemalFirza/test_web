package Test_Web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8081",
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
