package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tedsuo/rata"
)

func newIndexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "index")
	})
}

func newTestsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "tests")
	})
}

func main() {

	routes := rata.Routes{
		{Name: "get_index", Method: "GET", Path: "/"},
		{Name: "get_tests", Method: "GET", Path: "/tests"},
	}

	handlers := map[string]http.Handler{
		"get_index": newIndexHandler(),
		"get_tests": newTestsHandler(),
	}

	router, err := rata.NewRouter(routes, handlers)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())

}
