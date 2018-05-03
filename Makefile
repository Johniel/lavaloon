.PHONY: build
build:
	go build  -o a.out ./cmd/lavaloon

.PHONY: test
test:
	go test ./lavaloon
