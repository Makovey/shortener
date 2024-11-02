BINARY_NAME=shortener

remote_test:
	go build -o cmd/shortener/${BINARY_NAME} cmd/shortener/*.go
	./shortenertest -test.v -test.run=^TestIteration10$$ -source-path=. -binary-path=cmd/shortener/${BINARY_NAME} -database-dsn=postgres://admin:admin@localhost:5432/postgres