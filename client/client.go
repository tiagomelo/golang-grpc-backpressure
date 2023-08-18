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
	"net/http"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tiagomelo/golang-grpc-backpressure/api/proto/gen/stockservice"
	"github.com/tiagomelo/golang-grpc-backpressure/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	receivedUpdatesCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "stock_updates_received_total",
		Help: "The total number of stock updates received by the client",
	})
)

type options struct {
	RandomProcessingTime bool `short:"r" description:"Enable random processing time"`
}

// processStockUpdate simulates the processing of a stock update by sleeping a random time between 1 and 3 seconds.
func processStockUpdate(logger *log.Logger, update *stockservice.StockUpdate, randomProcessingTime bool) {
	if randomProcessingTime {
		const (
			sleepMin = 1
			sleepMax = 3
		)
		seed := time.Now().UnixNano()
		r := rand.New(rand.NewSource(seed))
		duration := time.Duration(r.Intn(sleepMax)+sleepMin) * time.Second
		time.Sleep(duration)
	}
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
func receiveStockUpdates(ctx context.Context, logger *log.Logger, client stockservice.StockServiceClient, randomProcessingTime bool) error {
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
		processStockUpdate(logger, update, randomProcessingTime)
		receivedUpdatesCounter.Inc()
	}
	return nil
}

func metricsHandler() http.Handler {
	return promhttp.Handler()
}

func metricsServer(serverPort int) {
	port := fmt.Sprintf(":%d", serverPort)
	http.Handle("/metrics", metricsHandler())
	log.Fatal(http.ListenAndServe(port, nil))
}

func run(logger *log.Logger, randomProcessingTime bool) error {
	logger.Println("main: initializing gRPC client")
	defer logger.Println("main: Completed")
	cfg, err := config.Read()
	if err != nil {
		return errors.Wrap(err, "reading config")
	}
	ctx := context.Background()
	const stockServiceHost = "localhost:4444"
	conn, err := grpc.Dial(stockServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to dial server:", err)
		os.Exit(1)
	}
	defer conn.Close()
	go metricsServer(cfg.StockServiceClientMetricsServerPort)
	client := stockservice.NewStockServiceClient(conn)
	if err := receiveStockUpdates(ctx, logger, client, randomProcessingTime); err != nil {
		return errors.Wrap(err, "receiving stock updates")
	}
	return nil
}

func main() {
	var opts options
	flags.Parse(&opts)
	logger := log.New(os.Stdout, "STOCK SERVICE CLIENT : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := run(logger, opts.RandomProcessingTime); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
