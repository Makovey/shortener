BINARY_NAME=shortener

remote_test:
	go build -o cmd/shortener/${BINARY_NAME} cmd/shortener/*.go
	./shortenertest -test.v -test.run=^TestIteration1$$ -binary-path=cmd/shortener/${BINARY_NAME}