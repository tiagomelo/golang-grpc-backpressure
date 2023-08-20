// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
//
// Package main contains the server-side implementation for a gRPC service that provides stock updates.
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tiagomelo/golang-grpc-backpressure/config"
	"github.com/tiagomelo/golang-grpc-backpressure/server"
)

// options struct holds command line flags configurations.
type options struct {
	InitialDelay int `short:"i" description:"Initial delay before sending stock updates" default:"0"`
}

// metricsHandler returns an HTTP handler for exposing Prometheus metrics.
func metricsHandler() http.Handler {
	return promhttp.Handler()
}

// metricsServer starts an HTTP server on a specified port to expose Prometheus metrics.
func metricsServer(cfg *config.Config) {
	port := fmt.Sprintf(":%d", cfg.StockServiceMetricsServerPort)
	http.Handle("/metrics", metricsHandler())
	log.Fatal(http.ListenAndServe(port, nil))
}

func run(logger *log.Logger, initialDelay int) error {
	logger.Println("main: initializing gRPC server")
	defer logger.Println("main: Completed")

	// Reading config.
	cfg, err := config.Read()
	if err != nil {
		return errors.Wrap(err, "reading config")
	}

	// Setting up a TCP listener on the specified port.
	port := fmt.Sprintf(":%d", cfg.GrpcServerPort)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "tcp listening")
	}

	// Initialize the gRPC server with the specified initial delay.
	srv := server.New(initialDelay)

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.

	serverErrors := make(chan error, 1)

	// Start the metrics server.
	go metricsServer(cfg)

	// Start the service listening for requests.
	go func() {
		logger.Printf("main: gRPC server listening on %s", port)
		serverErrors <- srv.GrpcSrv.Serve(lis)
	}()
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-shutdown:
		logger.Println("main: received signal for shutdown: ", sig)
		srv.GrpcSrv.Stop()
	}

	return nil
}

func main() {
	var opts options
	flags.Parse(&opts)
	logger := log.New(os.Stdout, "STOCK SERVICE SERVER : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	if err := run(logger, opts.InitialDelay); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
