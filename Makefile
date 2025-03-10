include .env

BINARY_NAME=shortener

SHELL := /bin/bash
LOCAL_MIGRATION_DIR=./internal/db/migrations
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

install-deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/golang/mock/mockgen@v1.6.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

mig-s:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

mig-u:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

mig-d:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

remote_all_tests:
	go build -o cmd/shortener/${BINARY_NAME} cmd/shortener/*.go;
	@for i in $(shell seq 1 24); do \
		./shortenertest -test.v -test.run=^TestIteration$$i$$ \
		-source-path=. \
		-server-port=8080 \
		-binary-path=cmd/shortener/${BINARY_NAME} \
		-file-storage-path=./urls.txt \
		-database-dsn=postgres://admin:admin@localhost:5432/postgres; \
	done

remote_current_iter:
	go build -o cmd/shortener/${BINARY_NAME} cmd/shortener/*.go
	./shortenertest -test.v -test.run=^TestIteration24$$ \
	-source-path=. \
	-binary-path=cmd/shortener/${BINARY_NAME} \
	-database-dsn=postgres://admin:admin@localhost:5432/postgres

lint:
	go build -o linter cmd/staticlint/main.go;
	./linter ./...

generate-service_info-api:
	mkdir -p internal/generated/service_info/
	protoc --proto_path api/service_info \
	--go_out=internal/generated/service_info/ --go_opt=paths=source_relative \
	--go-grpc_out=internal/generated/service_info/ --go-grpc_opt=paths=source_relative \
	api/service_info/service_info.proto

generate-shortener-api:
	mkdir -p internal/generated/shortener/
	protoc --proto_path api/shortener \
	--go_out=internal/generated/shortener/ --go_opt=paths=source_relative \
	--go-grpc_out=internal/generated/shortener/ --go-grpc_opt=paths=source_relative \
	api/shortener/shortener.proto