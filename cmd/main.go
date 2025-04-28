package main

import (
	"fmt"
	"net/http"
)

func main() {
	var err error
	//TO DO: log
	//TO DO: routing
	//TO DO: handlers

	mux := http.NewServeMux()
	mux.HandleFunc("/", mainPage)

	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		fmt.Printf("server start err%+v\n", err)
	}

}

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hrelloo"))
	fmt.Println(r.RequestURI)
}
