package main

import (
	"encoding/json"
	"net/http"
)

// JSON data representation
type Products struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	// 1. create multiplexer route
	mux := http.NewServeMux()

	// 3. create endpoint
	mux.HandleFunc("GET /products", listProduct)
	mux.HandleFunc("POST /products", createProduct)
	mux.HandleFunc("PUT /products/{id}", updateProduct)
	mux.HandleFunc("DELETE /products/{id}", deleteProduct)

	// 4. create server
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	server.ListenAndServe()
}

var database = map[int]Products{}

// 2. create func handler
func listProduct(w http.ResponseWriter, r *http.Request) {
	var products []Products
	for _, v := range database {
		products = append(products, v)
	}

	data, err := json.Marshal(products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("Terjadi Kesalahan"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

}

func createProduct(w http.ResponseWriter, r *http.Request) {

}

func updateProduct(w http.ResponseWriter, r *http.Request) {

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {

}
