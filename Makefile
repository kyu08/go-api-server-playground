# =========================================
# 開発環境構築
# =========================================
dev-tools:
	go install github.com/k1LoW/runn/cmd/runn@latest
	go install github.com/izumin5210/cgt@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	echo "--------------------------------------------------"
	echo "⚠️protoc, golangci-lint, sqlfluffは別途installしてください。"
	echo "--------------------------------------------------"

# =========================================
# 自動生成系
# =========================================
gen-sqlc: 
	sqlc generate

gen-all: gen-sqlc

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

lint-go:
	golangci-lint run -c ./.golangci.yaml --fix --allow-parallel-runners --tests ./...

build:
	go build ./...

lint-sql:
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

.PHONY: dev-tools gen-sqlc gen-all test test-e2e-with-refresh test-e2e-with-refresh-server test-e2e lint-go build handler-list health-check lint-sql container-up mysql-cli db-log container-stop container-restart
