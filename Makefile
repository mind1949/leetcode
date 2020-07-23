fmt:
	gofmt -w -s .
test:
	go test algs/*.go -cover -v
bench:
	go test algs/*.go -bench=. -run=NONE