package data

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var mtx sync.Mutex

var stockConfigs = map[string]bool{
	"AAPL": true,
	"AMZN": true,
	"GOOG": true,
	"META": true,
	"MSFT": true,
	"NFLX": true,
}

var stockPrices = map[string][]StockPrice{}

type StockPrice struct {
	Code      string
	Price     int64
	Timestamp time.Time
}

func init() {
	for code, isEnabled := range stockConfigs {
		if !isEnabled {
			continue
		}

		log.Printf("Stock %s is enabled", code)
		go updateStock(code)
	}
}

func updateStock(code string) {
	for {
		mtx.Lock()

		if !stockConfigs[code] {
			mtx.Unlock()
			break
		}

		time.Sleep(1 * time.Second)

		current, exist := stockPrices[code]
		if !exist {
			earlyPrice := 10000
			timeCreated := time.Now()

			stockPrices[code] = []StockPrice{
				{
					Code:      code,
					Price:     int64(earlyPrice),
					Timestamp: timeCreated,
				},
			}

			log.Printf("Stock %s price is %d at %s", code, earlyPrice, timeCreated)
			mtx.Unlock()

			continue
		}

		lastItem := current[len(stockPrices[code])-1]
		price := randomizePrice(lastItem.Price)

		stockPrices[code] = append(stockPrices[code], StockPrice{
			Code:      code,
			Price:     price,
			Timestamp: time.Now(),
		})

		log.Printf("Stock %s price is %d at %s", code, price, time.Now())

		mtx.Unlock()
	}
}

func randomizePrice(price int64) int64 {
	operation := rand.Intn(2)

	amount := rand.Int63n(100)

	if operation == 0 {
		return price - amount
	}

	return price + amount
}

func ToggleStock(code string, isEnabled bool) {
	mtx.Lock()

	if isEnabled == stockConfigs[code] {
		mtx.Unlock()
		return
	}

	log.Printf("Change stock %s status to %t", code, isEnabled)

	if !stockConfigs[code] && isEnabled {
		stockConfigs[code] = true

		mtx.Unlock()
		go updateStock(code)

		return
	}

	stockConfigs[code] = false

	mtx.Unlock()
}

func GetStockConfig() map[string]bool {
	return stockConfigs
}

func GetStockPrice(code string) []StockPrice {
	return stockPrices[code]
}
