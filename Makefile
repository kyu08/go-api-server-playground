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
# サーバー起動中(spanemuboostがエミュレータを起動中)にgen-yo-localを使用
gen-yo-local:
	@echo "Generating code from Spanner Emulator schema..."
	yo test-project test-instance test-database \
		-o internal/infrastructure/database/yo \
		-p yo

gen-all: gen-proto

# =========================================
# アプリケーションの起動、デバッグなど
# =========================================
run:
	go run cmd/server/main.go

test:
	go test -v ./... | cgt

# e2eテスト実行（サーバー起動→テスト→終了を一括実行）
test-e2e:
	@kill $$(lsof -ti:8080) 2>/dev/null || true
	@docker ps -q --filter "ancestor=gcr.io/cloud-spanner-emulator/emulator" | xargs -r docker stop 2>/dev/null || true
	@echo "Starting server..."
	@go run cmd/server/main.go & SERVER_PID=$$!; \
	echo "Waiting for server to be ready..."; \
	for i in $$(seq 1 60); do \
		if grpcurl -plaintext localhost:8080 twitter.TwitterService.Health >/dev/null 2>&1; then \
			echo "Server is ready!"; \
			break; \
		fi; \
		if [ $$i -eq 60 ]; then \
			echo "Server failed to start"; \
			kill $$(lsof -ti:8080) 2>/dev/null || true; \
			exit 1; \
		fi; \
		sleep 1; \
	done; \
	echo "Running e2e tests..."; \
	docker compose run --rm e2e; \
	TEST_EXIT_CODE=$$?; \
	kill $$(lsof -ti:8080) 2>/dev/null || true; \
	exit $$TEST_EXIT_CODE

lint-go:
	golangci-lint run -c ./.golangci.yaml --fix --allow-parallel-runners --tests ./...

build:
	go build ./...

handler-list:
	grpcurl -plaintext localhost:8080 list twitter.TwitterService

health-check:
	grpcurl -plaintext localhost:8080 twitter.TwitterService.Health

.PHONY: dev-tools gen-proto gen-yo-local gen-all run test test-e2e lint-go build handler-list health-check
