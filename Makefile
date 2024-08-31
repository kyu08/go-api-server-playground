# =========================================
# 開発環境構築
# =========================================
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
gen-proto: 
	cd pkg && protoc --go_out=./api --go_opt=paths=source_relative \
	--go-grpc_out=./api --go-grpc_opt=paths=source_relative \
	*.proto

gen-sqlc: 
	sqlc generate

gen-all: gen-proto gen-sqlc

# =========================================
# アプリケーションの起動、デバッグなど
# =========================================
test: 
	go test -v ./... | cgt

# すべてのコンテナをビルドしなおしてからE2Eを実行
test-e2e-with-refresh: container-up test-e2e

# serverコンテナだけビルドしなおしてからE2Eを実行
test-e2e-with-refresh-server: 
	docker compose build server && docker compose start server && make test-e2e

test-e2e: 
	docker compose run e2e

lint: 
	golangci-lint run -c ./.golangci.yaml --fix --allow-parallel-runners --tests ./...

build:
	go build ./...

handler-list:
	grpcurl -plaintext localhost:8080 list twitter.TwitterService

health-check:
	grpcurl -plaintext localhost:8080 twitter.TwitterService.Health

format-sql:
	sqlfluff format internal/infrastructure/database; sqlfluff fix --FIX-EVEN-UNPARSABLE internal/infrastructure/database; sqlfluff lint internal/infrastructure/database

# =========================================
# コンテナ関連
# =========================================
container-watch: # serverコンテナの変更を検知する度にビルドする
	docker compose watch

container-up:
	docker compose up -d --build --renew-anon-volumes --force-recreate --remove-orphans

mysql-cli:
	docker compose run mysql-cli

db-log:
	docker logs --tail 50 --follow --timestamps

container-stop:
	docker compose stop

container-restart:
	make container-stop
	docker container prune && docker volume rm $(docker volume ls -q)
	make container-up

.PHONY: dev-tools gen-proto gen-sqlc gen-all test test-e2e-with-refresh test-e2e-with-refresh-server test-e2e lint build handler-list health-check format-sql container-up mysql-cli db-log container-stop container-restart
