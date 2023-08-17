// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/tiagomelo/golang-grpc-backpressure/api/proto/gen/stockservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// processStockUpdate simulates the processing of a stock update by sleeping a random time between 1 and 3 seconds.
func processStockUpdate(logger *log.Logger, update *stockservice.StockUpdate) {
	const (
		sleepMin = 1
		sleepMax = 3
	)
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	duration := time.Duration(r.Intn(sleepMax)+sleepMin) * time.Second
	time.Sleep(duration)
	logger.Println(fmt.Sprintf(`ticker:"%s" price:%.2f change:%.2f changePercent:%.2f volume:%d openPrice:%.2f highPrice:%.2f lowPrice:%.2f marketCap:%d timestamp:"%s"`,
		update.Ticker,
		update.Price,
		update.Change,
		update.ChangePercent,
		update.Volume,
		update.OpenPrice,
		update.HighPrice,
		update.LowPrice,
		update.MarketCap,
		update.Timestamp,
	))
}

// receiveStockUpdates establishes a stream with the stock service to receive stock updates.
func receiveStockUpdates(ctx context.Context, logger *log.Logger, client stockservice.StockServiceClient) error {
	stream, err := client.GetUpdates(ctx, &stockservice.EmptyRequest{})
	if err != nil {
		return errors.Wrap(err, "opening stream")
	}
	for {
		update, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "receiving update")
		}
		processStockUpdate(logger, update)
	}
	return nil
}

func main() {
	ctx := context.Background()
	logger := log.New(os.Stdout, "STOCK SERVICE CLIENT : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	const stockServiceHost = "localhost:4444"
	conn, err := grpc.Dial(stockServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to dial server:", err)
		os.Exit(1)
	}
	defer conn.Close()
	client := stockservice.NewStockServiceClient(conn)
	if err := receiveStockUpdates(ctx, logger, client); err != nil {
		fmt.Println("Failed to receive stock updates:", err)
		os.Exit(1)
	}
}
