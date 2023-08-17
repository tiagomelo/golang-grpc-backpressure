// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package stock

import (
	"math/rand"
	"time"

	"github.com/tiagomelo/golang-grpc-backpressure/api/proto/gen/stockservice"
)

const (
	initialPrice          = 150.0
	priceFluctuationRange = 5.0
)

// stock represents a simple stock model.
type stock struct {
	ticker        string  // The stock's ticker symbol.
	currentPrice  float64 // The stock's current price.
	previousPrice float64 // The stock's price before the last update.
}

// New initializes and returns a new stock instance.
func New(ticker string, startingPrice float64) *stock {
	return &stock{
		ticker:        ticker,
		currentPrice:  startingPrice,
		previousPrice: startingPrice,
	}
}

// RandomUpdate generates a random stock update based on the stock's current price.
func (s *stock) RandomUpdate() *stockservice.StockUpdate {
	change := (rand.Float64() * priceFluctuationRange) - (priceFluctuationRange / 2)
	s.previousPrice = s.currentPrice
	s.currentPrice += change
	update := &stockservice.StockUpdate{
		Ticker:        s.ticker,
		Price:         s.currentPrice,
		Change:        change,
		ChangePercent: (change / s.previousPrice) * 100,
		Volume:        int64(rand.Intn(10000)),
		OpenPrice:     initialPrice,
		HighPrice:     s.currentPrice + rand.Float64()*2,
		LowPrice:      s.currentPrice - rand.Float64()*2,
		MarketCap:     int64(s.currentPrice * float64(rand.Intn(1000000))),
		Timestamp:     time.Now().Format(time.RFC3339),
	}
	return update
}
