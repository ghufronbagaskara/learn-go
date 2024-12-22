package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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



var database = map[int]Products{
	
}

// id for enumerate for database
var lastID = 0

// 2. create func handler
func listProduct(w http.ResponseWriter, r *http.Request) {
	// slice for response
	var products []Products

	// iteration to add product from map to product slice 
	for _, v := range database {
		products = append(products, v)
	}

	// change it to json
	data, err := json.Marshal(products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("Terjadi kesalahan"))
	}

	// add json to it
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

}

func createProduct(w http.ResponseWriter, r *http.Request) {
	
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Kesalahan dalam request"))
	}
	
	var products Products
	err = json.Unmarshal(bodyByte, &products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Kesalahan dalam request"))
	}

	// increment ID for input products
	lastID++

	products.ID = lastID
	
	// add data to database
	database[products.ID] = products
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("Request berhasil di proses dan ditambahkan"))
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	// read request id -> /{id}
	productID := r.PathValue("id")
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("Terjadi kesalahan"))
	}
	
	
	// read body r (request)
	bodyByte, err := io.ReadAll(r.Body) 
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Kesalahan dalam request"))
	}

	var products Products
	err = json.Unmarshal(bodyByte, &products) // unmarshal help change json to data class we had, which in this case Products struct
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Kesalahan dalam request"))
	}

	
	// assign prouduct ID from URL to ensure it is not replaced by the ID in the request body
	products.ID = productIDInt

	// update products with id 
	database[productIDInt] = products
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}


func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// read product id 
	productID := r.PathValue("id")

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Kesalahan dalam request"))
	}

	delete(database, productIDInt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}
