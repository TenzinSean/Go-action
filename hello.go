package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"github.com/julienschmidt/httprouter"
        "example.com/user/hello/morestring"
        "github.com/google/go-cmp/cmp"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
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

        mux := httprouter.New()
        mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr: "127.0.0.1:8080",
                Handler: mux,
	}
	//http.HandleFunc("/hello", hello)
	server.ListenAndServe()
        fmt.Println(morestrings.ReverseRunes("!oG, olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
