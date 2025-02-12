package main

import (
	"fmt"
	"sync"
	"time"
)


func main()  {
	// handle order
	order := Order{
		ID: 1,
		isValid: false,
	}
	// validate
	// distribute validation -> fanOut CH
	fanOutCH := make(chan func(), 4)
	// wait all validation task -> fanIn CH 
	fanInCH := make(chan bool, 4)

	var wg sync.WaitGroup

	// fanOutCH: distribute validation task
	wg.Add(4)
	go func() {
		fanOutCH <- func()  { validatePaymentStatus(&order, &wg, fanInCH) }
		fanOutCH <- func()  { validateSeller(&order, &wg, fanInCH) }
		fanOutCH <- func()  { validateStock(&order, &wg, fanInCH) }
		fanOutCH <- func()  { validateShippingAddress(&order, &wg, fanInCH) }
		close(fanOutCH)
	}()

	// fanInCH: collect result from all validation worker
	go func() {
		wg.Wait()
		close(fanInCH)
	}()

	// start worker 
	for f := range fanOutCH {
		go f()
	}


	order.isValid = true
	for result := range fanInCH {
		if !result {
			order.isValid = false
		}
	}
	

	if order.isValid {
		fmt.Printf("order isn't valid, cant be processed\n")
	} else {
		fmt.Printf("order is valid")
	}
	
}

type Order struct {
	ID int 
	isValid bool
}


func validatePaymentStatus(order *Order, wg *sync.WaitGroup, fanInCH chan<- bool) {
	defer wg.Done()
	
	time.Sleep(time.Millisecond * 200)
	fmt.Println("payment status validated: passed")
	fanInCH <- true
}

func validateSeller(order *Order, wg *sync.WaitGroup, fanInCH chan<- bool){
	defer wg.Done()
	
	time.Sleep(time.Millisecond * 150)
	fmt.Println("seller validattion: passed")
	fanInCH <- true
}

func validateStock(order *Order, wg *sync.WaitGroup, fanInCH chan<- bool){
	defer wg.Done()
	
	time.Sleep(time.Millisecond * 200)
	fmt.Println("stock validated: passed")
	fanInCH <- true
}

func validateShippingAddress(order *Order, wg *sync.WaitGroup, fanInCH chan<- bool) {
	defer wg.Done()
	
	time.Sleep(time.Millisecond * 50)
	fmt.Println("shipping address validated: passed")
	fanInCH <- true
}