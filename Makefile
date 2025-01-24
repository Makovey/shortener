include .env

BINARY_NAME=shortener

SHELL := /bin/bash
LOCAL_MIGRATION_DIR=./internal/db/migrations
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

install-deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/golang/mock/mockgen@v1.6.0

mig-s:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

mig-u:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

mig-d:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

remote_all_tests:
	go build -o cmd/shortener/${BINARY_NAME} cmd/shortener/*.go;
	@for i in $(shell seq 1 16); do \
		./shortenertest -test.v -test.run=^TestIteration$$i$$ \
		-source-path=. \
		-server-port=8080 \
		-binary-path=cmd/shortener/${BINARY_NAME} \
		-file-storage-path=./urls.txt \
		-database-dsn=postgres://admin:admin@localhost:5432/postgres; \
	done

remote_current_iter:
	go build -o cmd/shortener/${BINARY_NAME} cmd/shortener/*.go
	./shortenertest -test.v -test.run=^TestIteration15$$ \
	-source-path=. \
	-binary-path=cmd/shortener/${BINARY_NAME} \
	-database-dsn=postgres://admin:admin@localhost:5432/postgres