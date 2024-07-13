fmt:
	goimports -w . .
	gofmt -w -s .
	cargo fmt
test:
	go test algs/*.go -cover -v
bench:
	go test algs/*.go -bench=. -run=NONE
