package requester

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func newMux() http.Handler {
	m := http.NewServeMux()

	v1 := v1Handler()

	m.Handle("/v1", http.StripPrefix("/v1", v1))
	return m
}

func newHttp() *http.Server {
	return &http.Server{
		Addr:         "localhost:8080",
		Handler:      newMux(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func InitHttpServer() {
	s := newHttp()

	if err := s.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
