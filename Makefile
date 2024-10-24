BINARY_NAME=shortener

remote_test:
	go build -o cmd/shortener/${BINARY_NAME} cmd/shortener/*.go
	./shortenertest -test.v -test.run=^TestIteration7$$ -source-path=cmd/shortener/${BINARY_NAME} -binary-path=cmd/shortener/${BINARY_NAME} -server-port=8999