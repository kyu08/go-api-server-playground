.PHONY: gen-proto
gen-proto:
	cd api && protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
	*.proto

.PHONY: run
run:
	go run ./cmd/server/main.go

.PHONY: build
build:
	go build ./...

.PHONY: resolver-list
resolver-list:
	grpcurl -plaintext localhost:8080 list

.PHONY: test-req # あとで消す
test-req:
	grpcurl -plaintext localhost:8080  twitter.TwitterService.Health
