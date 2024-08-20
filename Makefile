# =========================================
# 開発環境構築
# =========================================
.PHONY: dev-tools
dev-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/k1LoW/runn/cmd/runn@latest
	go install github.com/izumin5210/cgt@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	echo "--------------------------------------------------"
	echo "⚠️protoc, golangci-lint, sqlfluffは別途installしてください。"
	echo "--------------------------------------------------"

# =========================================
# 自動生成系
# =========================================
.PHONY: gen-proto
gen-proto: 
	cd api && protoc --go_out=../pkg/api --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/api --go-grpc_opt=paths=source_relative \
	*.proto

.PHONY: gen-sqlc
gen-sqlc: 
	sqlc generate

.PHONY: gen-all
gen-all: gen-proto gen-sqlc

# =========================================
# アプリケーションの起動、デバッグなど
# =========================================
.PHONY: test
test: 
	go test -v ./... | cgt

.PHONY: test-e2e-with-refresh # goコードの変更後に実行したいケース
test-e2e-with-refresh: container-up test-e2e

.PHONY: test-e2e
test-e2e: 
	docker compose run e2e

.PHONY: lint
lint: 
	golangci-lint run -c ./.golangci.yaml --fix --allow-parallel-runners --tests ./...

.PHONY: build
build:
	go build ./...

.PHONY: handler-list
handler-list:
	grpcurl -plaintext localhost:8080 list twitter.TwitterService

.PHONY: health-check
health-check:
	grpcurl -plaintext localhost:8080 twitter.TwitterService.Health

.PHONY: format-sql
format-sql:
	sqlfluff format internal/database; sqlfluff fix --FIX-EVEN-UNPARSABLE internal/database; sqlfluff lint internal/database

# =========================================
# コンテナ関連
# =========================================
.PHONY: container-up
container-up:
	docker compose up -d --build --renew-anon-volumes --force-recreate --remove-orphans

.PHONY: mysql-cli
mysql-cli:
	docker compose run mysql-cli

.PHONY: db-log # とはいえDocker Desktopでみた方がわかりやすそうではある
db-log:
	docker logs --tail 50 --follow --timestamps

.PHONY: container-stop
container-stop:
	docker compose stop

.PHONY: container-restart
container-restart:
	make container-stop
	docker container prune && docker volume rm $(docker volume ls -q)
	make container-up
