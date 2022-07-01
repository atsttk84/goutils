.PHONY: fmt test

fmt:
	gofmt -l -s -w .

test:
	go test -count 1 -v ./...
