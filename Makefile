include .env
export

POSTGRES_SETUP=user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) host=$(POSTGRES_DB_HOST) port=5432 sslmode=disable

MIGRATION_FOLDER=./migrations
PROTO_FOLDER=./internal/pkg/pb

migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up

migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" down


build:
	go build -o main ./cmd

run:
	make build
	make migration-up
	./main

generate:
	rmdir /s/q "$(PROTO_FOLDER)"
	mkdir "$(PROTO_FOLDER)"
	protoc \
		--proto_path="./proto" \
		--go_out="$(PROTO_FOLDER)" \
		--go-grpc_out="$(PROTO_FOLDER)" \
		"./proto/*.proto"
