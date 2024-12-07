BINARY_NAME=shortener

remote_all_tests:
	go build -o cmd/shortener/${BINARY_NAME} cmd/shortener/*.go;
	@for i in $(shell seq 1 15); do \
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