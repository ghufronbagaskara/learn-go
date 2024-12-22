// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// // struct with json struct text
// type FactResponse struct {
// 	Text string `json:"text"`
// 	Type string `json:"type"`
// }

// func main() {
// 	// 1. buat request
// 	req, err := http.NewRequest("GET", "https://cat-fact.herokuapp.com/facts/591f98803b90f7150a19c229", nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(0)
// 	}

// 	// 2. init client
// 	client := http.Client{}

// 	// 3. call req from client
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(0)
// 	}
// 	fmt.Println(res)

// 	// close response body
// 	defer res.Body.Close()

// 	// 4. read response body
// 	resBody, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(0)
// 	}

// 	// 6. convert to FactResponse data type
// 	var factResponse FactResponse
// 	json.Unmarshal(resBody, &factResponse)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(0)
// 	}

// 	fmt.Println("text : ", factResponse.Text)
// 	fmt.Println("type : ", factResponse.Type)
// }