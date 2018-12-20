package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/alice"
)

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

type headerSetter struct {
	key, val string
	handler  http.Handler
}

func (hs headerSetter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(hs.key, hs.val)
	hs.handler.ServeHTTP(w, r)
}

func newHeaderSetter(key, val string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		log.Printf("In Header setter .....")
		return headerSetter{key, val, h}
	}
}

func main() {
	h := http.NewServeMux()

	h.HandleFunc("/nks", myFun)

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "hit url: /nks")
	})

	chain := alice.New(
		newHeaderSetter("X-FOO", "BAR"),
		newHeaderSetter("X-BAZ", "BUZ"),
		logger,
	).Then(h)

	err := http.ListenAndServe(":9999", chain)
	log.Fatal(err)
}

func myFun(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Nandeshwar"))
}
