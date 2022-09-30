.PHONY: init
init:
	docker-compose exec goapp sh /init.sh

.PHONY: fmt
fmt:
	gofmt -l -s -w .

.PHONY: test
test:
	go mod tidy
	go test -count 1 -v ./...
