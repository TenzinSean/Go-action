package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	//"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called -" + name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)
	server.ListenAndServe()
}
