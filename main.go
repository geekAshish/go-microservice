package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/geekashish/go-microservice/handler"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handler.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}
