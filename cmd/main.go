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
	mux.HandleFunc("/", withFullRequestLoging(mainPage))

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
		log.Println("duration: ", duration)

	}

}

func withFullRequestLoging(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Printf("\nr.Body : ", r.Body,
			"\nr.ContentLength : ", r.ContentLength,
			"\nr.Form: ", r.Form,
			"\nr.GetBody: ", r.GetBody,
			"\nr.Header: ", r.Header,
			"\nr.Host: ", r.Host,
			"\nr.Method : ", r.Method,
			"\nr.MultipartForm : ", r.MultipartForm,
			"\nr.Pattern : ", r.Pattern,
			"\nr.PostForm: ", r.PostForm,
			"\nr.Proto: ", r.Proto,
			"\nr.ProtoMajor: ", r.ProtoMajor,
			"\nr.ProtoMinor: ", r.ProtoMinor,
			"\nr.RemoteAddr: ", r.RemoteAddr,
			"\nr.RequestURI: ", r.RequestURI,
			"\nr.Response: ", r.Response,
			"\nr.TLS: ", r.TLS,
			"\nr.Trailer: ", r.Trailer,
			"\nr.Trailer: ", r.Trailer,
			"\nr.URL: ", r.URL)
	}
}
func withFullResponseLoging(h http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w,r)
		
		log.Printf(" request loging: ",
		"\n w.Header func: ",w.Header())
		log.Printf("\n w.Write func: ")
		w.Write([]byte("Some slice Byte fn write"))
		log.Println("\n w.WriteHeader func: ")
		w.WriteHeader(1234)

		log.Println("\n w.Header.Add func: ")
		w.Header().Add()

		log.Println("\n w.Header().Clone() func: ")
		w.Header().Clone()

		log.Println("\n w.Header().Del() func: ")
		w.Header().Del(),

	}


}

