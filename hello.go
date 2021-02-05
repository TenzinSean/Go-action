package main

import (
	"fmt"
	//"net/http"
	//"reflect"
	//"runtime"
	//"github.com/julienschmidt/httprouter"
        //"example.com/user/hello/morestring"
        //"github.com/google/go-cmp/cmp"
    //"golang.org/x/net/http2"
        "github.com/labstack/echo/v4"
       "github.com/labstack/echo/v4/middleware"
)

// type MyHandler struct {}

// func (h *MyHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!")
// }

// func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
// }

// func world(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello")
// }

// func log(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
// 		fmt.Println("Handler function called -" + name)
// 		h(w, r)
// 	}
// }


func main() {

        e := echo.New()
        e.Use(middleware.Logger())
        e.Use(middleware.Recover())

        e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})

        e.GET("/file", func(c echo.Context) error {
		return c.File("echo.svg")
	})


        e.Logger.Fatal(e.Start(":1232"))
	//server := http.Server{
	//	Addr: "127.0.0.1:8080",
	//}
	//http.HandleFunc("/Headers", headers)
	//server.ListenAndServe()

    //     http2.ConfigureServer(&server, &http2.Server{})
	// server.ListenAndServeTLS("cert.pem", "key.pem")
        //fmt.Println(morestrings.ReverseRunes("!oG, olleH"))
	//fmt.Println(cmp.Diff("Hello World", "Hello Go"))
         fmt.Println("hhiii")
}
