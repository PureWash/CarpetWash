# Include environment variables
# include .env

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

# Migrate database using environment variables from .env
.PHONY: migrate
migrate:
	migrate -source file://migrations -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable" up

# Database URL (local development)
DB_URL := postgres://postgres:123@localhost:5432/db?sslmode=disable

# Run migrations up
.PHONY: migrate-up
migrate-up:
	migrate -path migrations -database $(DB_URL) -verbose up

# Rollback migrations (down)
.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database $(DB_URL) -verbose down

# Force migrate to a specific version (in this case, 1)
.PHONY: migrate-force
migrate-force:
	migrate -path migrations -database $(DB_URL) -verbose force 1

# Create new migration file with a sequential name
.PHONY: migrate-file
migrate-file:
	migrate create -ext sql -dir migrations/ -seq carpet_wash

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
