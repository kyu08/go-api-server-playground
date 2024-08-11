.PHONY: gen-proto
gen-proto: tools-proto
	cd api && protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
	*.proto

.PHONY: gen-sqlc
gen-sqlc: tools-sqlc
	sqlc generate

.PHONY: run
run: tools-run
	air

.PHONY: test
test: tools-test
	go test -v ./... | cgt

.PHONY: test-e2e
test-e2e: tools-test-e2e
	runn run --grpc-no-tls e2e/**/*.yaml

.PHONY: lint
lint: tools-lint
	golangci-lint run -c ./.golangci.yaml --fix --tests ./...

.PHONY: build
build:
	go build ./...

.PHONY: start-db
start-db:
	docker compose up -d --build --renew-anon-volumes --force-recreate mysql

.PHONY: run-db-cli
run-db-cli:
	docker compose run cli

.PHONY: db-log
db-log:
	docker logs --tail 50 --follow --timestamps

.PHONY: container-stop
container-stop:
	docker compose stop

.PHONY: container-restart
container-restart:
	make container-stop
	docker volume rm $(docker volume ls -qf dangling=true)
	make start-db

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

.PHONY: tools-test-e2e
tools-test-e2e:
	@if ! which runn > /dev/null; then \
		echo "Please install k1LoW/runn"; \
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

.PHONY: tools-sqlc
tools-sqlc:
	@if ! which sqlc > /dev/null; then \
		echo "Please install sqlc"; \
	fi
