package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	println("starting http server ...")

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`pong`))
	})

	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("Server error: %s", err)
	println("stopped http server")
}
