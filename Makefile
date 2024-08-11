.PHONY: gen-proto
gen-proto: tools-proto
	cd api && protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
	*.proto

.PHONY: run
run: tools-run
	air

.PHONY: test
test: tools-test
	go test -v ./... | cgt

.PHONY: lint
lint: tools-lint
	golangci-lint run -c ./.golangci.yaml --fix --tests ./...

.PHONY: build
build:
	go build ./...

.PHONY: resolver-list
resolver-list:
	grpcurl -plaintext localhost:8080 list

.PHONY: test-req # TODO: E2Eを導入したら消す
test-req:
	grpcurl -plaintext localhost:8080 twitter.TwitterService.Health

.PHONY: tools-proto
tools-proto:
	@if ! which protoc > /dev/null; then \
		echo "Please install protoc"; \
	fi

.PHONY: tools-run
tools-run:
	@if ! which air > /dev/null; then \
		echo "Please install air-verse/air"; \
	fi

.PHONY: tools-test
tools-test:
	@if ! which cgt > /dev/null; then \
		echo "Please install izumin5210/cgt"; \
	fi

.PHONY: tools-lint
tools-lint:
	@if ! which golangci-lint > /dev/null; then \
		echo "Please install golangci-lint"; \
	fi
