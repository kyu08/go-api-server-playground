.PHONY: gen-proto
gen-proto:
	cd api && protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
	*.proto

.PHONY: test-req # あとで消す
test-req:
	grpcurl -plaintext -d '{"name": "kyu"}' localhost:8080  myapp.GreetingService.Hello
