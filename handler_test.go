package Test_Web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hallo ngab apa kabar ?")
	}

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welkam Tu Mai Cenel")
	})
	mux.HandleFunc("/saskrep", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "plis saskrep")
	})

	server := http.Server{
		Addr:    "localhost:8084",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
