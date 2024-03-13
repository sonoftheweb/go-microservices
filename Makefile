# Makefile located in go-ms/

# Variables
AUTH_SERVICE_PATH=./auth
USER_SERVICE_PATH=./user
GATEWAY_SERVICE_PATH=./gateway
GATEWAY_PROTO_DEST=$(GATEWAY_SERVICE_PATH)/

AUTH_PROTO_PATH=$(AUTH_SERVICE_PATH)/pb
AUTH_PROTO_FILE=$(AUTH_PROTO_PATH)/auth.proto

USER_PROTO_PATH=$(USER_SERVICE_PATH)/pb
USER_PROTO_FILE=$(USER_PROTO_PATH)/user.proto

# Build the auth service Docker image
build-auth:
	cd $(AUTH_SERVICE_PATH) && docker build -t auth .

# Build the auth service Docker image
build-user:
	cd $(USER_SERVICE_PATH) && docker build -t user .

# Generate gRPC code from the proto file for both auth service and API gateway
gen-proto:
	protoc --go_out=$(AUTH_SERVICE_PATH) --go-grpc_out=$(AUTH_SERVICE_PATH) $(AUTH_PROTO_FILE)
	protoc --proto_path=$(AUTH_PROTO_PATH) --go_out=$(GATEWAY_PROTO_DEST) --go-grpc_out=$(GATEWAY_PROTO_DEST) $(AUTH_PROTO_FILE)
	protoc --go_out=$(USER_SERVICE_PATH) --go-grpc_out=$(USER_SERVICE_PATH) $(USER_PROTO_FILE)
	protoc --proto_path=$(USER_PROTO_PATH) --go_out=$(GATEWAY_PROTO_DEST) --go-grpc_out=$(GATEWAY_PROTO_DEST) $(USER_PROTO_FILE)

# Docker Compose up
up:
	docker-compose up --build -d

# Docker Compose down
down:
	docker-compose down

# Target to run all necessary setups
setup: gen-proto build-auth build-user

.PHONY: build-auth build-user gen-proto up down setup
