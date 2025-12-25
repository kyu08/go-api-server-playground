# =========================================
# 開発環境構築
# =========================================
dev-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/k1LoW/runn/cmd/runn@latest
	go install github.com/izumin5210/cgt@latest
	go install go.mercari.io/yo@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	echo "--------------------------------------------------"
	echo "⚠️protoc, golangci-lintは別途installしてください。"
	echo "--------------------------------------------------"

# =========================================
# 自動生成系
# =========================================
gen-proto:
	cd pkg && protoc --go_out=./api --go_opt=paths=source_relative \
	--go-grpc_out=./api --go-grpc_opt=paths=source_relative \
	*.proto

# NOTE: yoによるコード生成はSpanner EmulatorまたはSpannerインスタンスに接続して行う
# ローカル開発時はgen-yo-localを使用
gen-yo-local:
	@echo "Generating code from Spanner Emulator schema..."
	yo test-project test-instance test-database \
		-o internal/infrastructure/database/yo \
		-p yo

gen-all: gen-proto

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

handler-list:
	grpcurl -plaintext localhost:8080 list twitter.TwitterService

health-check:
	grpcurl -plaintext localhost:8080 twitter.TwitterService.Health

# =========================================
# コンテナ関連
# =========================================
container-watch: # serverコンテナの変更を検知する度にビルドする
	docker compose watch

container-up:
	docker compose up -d --build --renew-anon-volumes --force-recreate --remove-orphans

spanner-logs:
	docker compose logs -f spanner

container-stop:
	docker compose stop

container-restart:
	make container-stop
	docker container prune -f && docker volume rm $$(docker volume ls -q) 2>/dev/null || true
	make container-up

.PHONY: dev-tools gen-proto gen-yo-local gen-all test test-e2e-with-refresh test-e2e-with-refresh-server test-e2e lint-go build handler-list health-check container-up spanner-logs container-stop container-restart
