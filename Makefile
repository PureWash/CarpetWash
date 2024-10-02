# Include environment variables
include .env
export $(shell sed 's/=.*//' .env)

# Get the current directory
CURRENT_DIR := $(shell pwd)

# Run the Go service
.PHONY: run
run:
	go run cmd/main.go

# Generate protobuf files
.PHONY: proto-gen
proto-gen:
	./scripts/gen_proto.sh

# Pull proto submodule
.PHONY: pull-proto-module
pull-proto-module:
	git submodule update --init --recursive

# Update proto submodule
.PHONY: update-proto-module
update-proto-module:
	git submodule update --remote --merge

# Initialize Swagger docs
.PHONY: swag-init
swag-init:
	swag init -g api/router.go -o api/docs

# Generate SQLC files
.PHONY: sqlc-generate
sqlc-generate:
	sqlc vet; sqlc generate
