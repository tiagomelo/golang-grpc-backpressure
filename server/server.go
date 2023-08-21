// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package server

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tiagomelo/golang-grpc-backpressure/api/proto/gen/stockservice"
	"github.com/tiagomelo/golang-grpc-backpressure/stock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// sentUpdatesCounter is a Prometheus metric to keep track of the number of sent stock updates.
var sentUpdatesCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "stock_updates_sent_total",
	Help: "The total number of stock updates sent by the server",
})

// server struct holds the gRPC server instance and implements the StockServiceServer interface.
type server struct {
	stockservice.UnimplementedStockServiceServer
	GrpcSrv      *grpc.Server
	initialDelay int
}

// New initializes and returns a new gRPC server with the StockService registered.
func New(initialDelay int) *server {
	grpcServer := grpc.NewServer()
	srv := &server{
		GrpcSrv:      grpcServer,
		initialDelay: initialDelay,
	}

	// Register the StockService with the gRPC server instance.
	stockservice.RegisterStockServiceServer(grpcServer, srv)

	// Register reflection service on gRPC server, useful for tools like `grpcurl`.
	reflection.Register(grpcServer)
	return srv
}

// GetUpdates streams stock updates to the client. It creates a stock with a starting price and sends
// random updates to the connected client every second.
func (s *server) GetUpdates(_ *stockservice.EmptyRequest, stream stockservice.StockService_GetUpdatesServer) error {
	const (
		ticker        = "AAPL"
		startingPrice = 150.0
	)
	stock := stock.New(ticker, startingPrice)
	time.Sleep(time.Duration(s.initialDelay) * time.Second)
	for {
		update := stock.RandomUpdate()
		if err := stream.Send(update); err != nil {
			return status.Error(codes.Unknown, "failed to send update to client: "+err.Error())
		}
		sentUpdatesCounter.Inc()
		time.Sleep(100 * time.Millisecond)
	}
}
