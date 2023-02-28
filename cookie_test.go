package Test_Web

import (
	"fmt"
	"net/http"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "Qemsky17"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "Success Creating Cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Qemsky17")
	if err != nil {
		fmt.Fprint(w, "You Have No Cookie, please set one")
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "Hello %s", name)
	}

}
func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)
	server := http.Server{
		Addr:    "localhost:8085",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
