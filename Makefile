# ==============================================================================
# Help

.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]\n"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==============================================================================
# Protofile compilation

.PHONY: proto
## proto: compile proto files
proto:
	@ rm -rf api/proto/gen/stockservice
	@ mkdir -p api/proto/gen/stockservice
	@ cd api/proto ; \
	protoc --go_out=gen/stockservice --go_opt=paths=source_relative --go-grpc_out=gen/stockservice --go-grpc_opt=paths=source_relative stockservice.proto


# ==============================================================================
# gRPC server execution

.PHONY: server
## server: runs gRPC server
server:
	@ go run cmd/main.go

.PHONY: client
## client: runs gRPC client
client:
	@ go run client/client.go