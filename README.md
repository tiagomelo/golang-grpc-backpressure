# golang-grpc-backpressure

A sample project showing how gRPC handles back pressure.

## Prerequisites

- [Docker](docker.com)
- [docker-compose](docs.docker.com/compose/)

## running it

### server 

```
$ make server
```

### client 

```
$ make client
```

This runs the client that process the stock update as soon as they arrive.

### busy client 

```
$ make client-random-processing-time
```

This runs the client that sleeps at random time before processing the stock update.

## Grafana

```
make obs
```

Then head out to `localhost:3000`. Password is defined in `GF_SECURITY_ADMIN_PASSWORD` env var in `.env` file.

The dashboard name is `sockUpdatesSentVersusStockUpdatesProcessed`.

To stop Grafana:

```
make obs-stop
```

This will delete both Grafana and Prometheus volumes as well.