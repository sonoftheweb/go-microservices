# Makefile located in go-ms/

# Variables
AUTH_SERVICE_PATH=./auth
GATEWAY_SERVICE_PATH=./gateway
PROTO_PATH=$(AUTH_SERVICE_PATH)/pb
PROTO_FILE=$(PROTO_PATH)/auth.proto
GATEWAY_PROTO_DEST=$(GATEWAY_SERVICE_PATH)/

# Build the auth service Docker image
build-auth:
	cd $(AUTH_SERVICE_PATH) && docker build -t auth .

# Generate gRPC code from the proto file for both auth service and API gateway
gen-proto:
	protoc --go_out=$(AUTH_SERVICE_PATH) --go-grpc_out=$(AUTH_SERVICE_PATH) $(PROTO_FILE)
	protoc --proto_path=$(PROTO_PATH) --go_out=$(GATEWAY_PROTO_DEST) --go-grpc_out=$(GATEWAY_PROTO_DEST) $(PROTO_FILE)

# Docker Compose up
up:
	docker-compose up --build -d

# Docker Compose down
down:
	docker-compose down

# Target to run all necessary setups
setup: gen-proto build-auth

.PHONY: build-auth gen-proto up down setup
