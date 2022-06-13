package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	//create variable router
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintln(writer, "Hello HttpRouter")
		fmt.Fprintln(writer, "Hello unyu")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}
	server.ListenAndServe()
}
