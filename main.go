package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tedsuo/rata"
)

const viewsDir = "public"

func newTestsHandler(c *Collector) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(c.CollectResults())
	})
}

func newServer(args []string) (*http.Server, error) {

	c, err := NewCollector(args[0])
	if err != nil {
		fmt.Errorf("Error creating test collector: %s\n", err)
	}
	log.Printf("Watching reports dir [%s]\n", c.ReportsDir)

	routes := rata.Routes{
		{Name: "get_index", Method: "GET", Path: "/"},
		{Name: "get_tests", Method: "GET", Path: "/tests"},
	}

	handlers := map[string]http.Handler{
		"get_index": http.FileServer(http.Dir(viewsDir)),
		"get_tests": newTestsHandler(c),
	}

	router, err := rata.NewRouter(routes, handlers)
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server, nil
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "usage: gawker [reports dir]\n")
		os.Exit(1)
	}

	s, err := newServer(args)
	if err != nil {
		fmt.Errorf("Error creating server: ", err)
		os.Exit(1)
	}
	log.Fatal(s.ListenAndServe())
}
