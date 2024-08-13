# =========================================
# 開発環境構築
# =========================================
.PHONY: dev-tools
dev-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/k1LoW/runn/cmd/runn@latest
	go install github.com/izumin5210/cgt@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	echo "--------------------------------------------------"
	echo "⚠️golangci-lintは別途installしてください。"
	echo "--------------------------------------------------"

# =========================================
# 自動生成系
# =========================================
.PHONY: gen-proto
gen-proto: 
	cd api && protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
	*.proto

.PHONY: gen-sqlc
gen-sqlc: 
	sqlc generate

# =========================================
# アプリケーションの起動、デバッグなど
# =========================================
.PHONY: test
test: 
	go test -v ./... | cgt

.PHONY: test-e2e # たいていコードの変更後に実行したいのでコンテナの再ビルドも含めてしまう
test-e2e: 
	make container-up && docker compose run e2e

.PHONY: lint
lint: 
	golangci-lint run -c ./.golangci.yaml --fix --tests ./...

.PHONY: build
build:
	go build ./...

.PHONY: handler-list
handler-list:
	grpcurl -plaintext localhost:8080 list twitter.TwitterService

.PHONY: health-check
health-check:
	grpcurl -plaintext localhost:8080 twitter.TwitterService.Health

# =========================================
# コンテナ関連
# =========================================
.PHONY: container-up
container-up:
	docker compose up -d --build --renew-anon-volumes --force-recreate --remove-orphans

.PHONY: mysql-cli
mysql-cli:
	docker compose run mysql-cli

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
	make container-up
