package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main()  {
	// simulate fetch order
	// simulate filtering order
	// simulate analysing order item  	
	// simulate storing the analysis result to big query

	orderCH := make(chan Order)
	filteredOrderCH := make(chan Item)
	analysisReportCH := make(chan AnalysisReport)



	// run stage of pipeline concurrently 
	go fetchOrder(orderCH)
	go filterOrder(orderCH, filteredOrderCH)
	go analyseOrder(filteredOrderCH, analysisReportCH)
	go storeAnalysisReport(analysisReportCH)

	select {}

}



type (
	Order struct {
		ID int
		Items []Item
	}

	Item struct {
		ItemID int
		Category string // to be filterred, specific for "digital" category
		ProviderName string // telkomsel, indosat, esia
		Price float64
	}

	AnalysisReport struct {
		Category string
		AveragePrice float64
		MinPrice float64
		MaxPrice float64
	}
)


func fetchOrder(orderCH chan <- Order){
	for i := 0; ; i++ {
		orderData := Order{
			ID: i,
			Items: []Item{
				{
					ItemID: i*10 + 2,
					Category: "digital",
					ProviderName: "telkomsel",
					Price: rand.Float64() * 10000,
				},
			},
		}

		time.Sleep(500 * time.Millisecond)
		orderCH <- orderData
		fmt.Println("fetched the order ")
		
	}
}


func filterOrder(orderCH <-chan Order, filteredOrderCH chan<- Item) {
	for order := range orderCH {
		for _, item := range order.Items {
			if item.Category == "digital" {
				time.Sleep(200 * time.Millisecond)	
				filteredOrderCH <- item
				fmt.Println("filtered order detail") 
			}

		}
	}
}


func analyseOrder(filteredOrderCH <-chan Item, analysisReportCH chan<- AnalysisReport){
	for item := range filteredOrderCH {
		// do analyse here
		result := AnalysisReport {
			Category: item.Category,
			MaxPrice: item.Price * 2, // to simplify the calculation
			MinPrice: item.Price, // to simplify the calculation
			AveragePrice: item.Price, // to simplify the calculation
		}

		analysisReportCH <- result
		fmt.Println("analyzed the order detail")
	}

}


func storeAnalysisReport(analysisReportCH <-chan AnalysisReport){
	for report := range analysisReportCH{
		// store the analysis report to big query
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("store the analysis report to BQ with result: %+v\n\n", report)
	}
}

