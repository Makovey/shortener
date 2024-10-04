BINARY_NAME=shortener

remote_test:
	go build -o cmd/shortener/${BINARY_NAME} cmd/shortener/*.go
	./shortenertest -test.v -test.run=^TestIteration5$$ -binary-path=cmd/shortener/${BINARY_NAME} -server-port=8999