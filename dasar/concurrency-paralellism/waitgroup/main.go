package main

import "sync"


func main() {
	fetchPricing()
}

func fetchPricing() {
	var wg sync.WaitGroup
	// start
	// fetch indosat, telkomsel, and xl with gorouting
	// done

	wg.Add(1)
	indosatPricingResult, indosatErr := FetchIndosatAPI(&wg)

	wg.Add(1)
	telkomselPricingResult, telkomselErr := FetchTelkomselAPI(&wg)

	wg.Add(1)
	xlPricingResult, xlError := FetchXlAPI(&wg)

	wg.Wait()

	// cache result to redis
	_ = indosatPricingResult
	_ = indosatErr
	_ = telkomselPricingResult
	_ = telkomselErr
	_ = xlPricingResult
	_ = xlError
	



}



func FetchIndosatAPI(wg *sync.WaitGroup) (data struct{}, err error) {
	defer wg.Done()

	// fetch indosat
	return
}

func FetchTelkomselAPI(wg *sync.WaitGroup) (data struct{}, err error) {
	defer wg.Done()

	// fetch telkomsel
	return
}

func FetchXlAPI(wg *sync.WaitGroup) (data struct{}, err error) {
	defer wg.Done()

	// fetch xl
	return
}
