.PHONY: gen-proto
gen-proto:
	cd api && protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
	*.proto

.PHONY: tools
tools:
	@go build -mod=mod -o ./_bin/air		github.com/air-verse/air
	@go build -mod=mod -o ./_bin/cgt		github.com/izumin5210/cgt

.PHONY: run
run:
	./_bin/air

.PHONY: test
test:
	go test -v ./... | ./_bin/cgt

.PHONY: build
build:
	go build ./...

.PHONY: resolver-list
resolver-list:
	grpcurl -plaintext localhost:8080 list

.PHONY: test-req # TODO: E2Eを導入したら消す
test-req:
	grpcurl -plaintext localhost:8080 twitter.TwitterService.Health
