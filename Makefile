.PHONY: test install

install:
	go get -t -v ./...

test: install
	diff -u <(echo -n) <(gofmt -d -s .)
	go test -cover -bench=. -v -race ./...