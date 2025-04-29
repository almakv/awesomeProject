package main

import (
	"awesomeProject/internal/logger"
	"log"
	"net/http"
	"time"
)

var logg *log.Logger
var err error

func main() {
	logg = logger.LogLoad()
	logg.SetPrefix("/cmd/main/main.go ")
	mux := http.NewServeMux()
	mux.HandleFunc("/", withLoging(mainPage))
	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		logg.Fatalln("server init error")
	}
}
func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hrelloo"))
	if r.Method == "GET" {
	}
}
func withLoging(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logg.SetPrefix("\n/cmd/main.go fn handler ")
		start := time.Now()

		h.ServeHTTP(w, r)

		duration := time.Since(start).Microseconds()
		logg.Println("\nmethod", r.Method, "\nt: ", duration, "millisec\n", "transEnc", r.TransferEncoding)
		logg.
	}
}
