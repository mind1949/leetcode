fmt:
	gofmt -w -s .
test:
	go test algs/*.go -cover -v