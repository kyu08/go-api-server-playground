FROM golang:1.23.0-alpine AS builder
WORKDIR /app

# モジュールキャッシュを効かせるために必要
COPY go.mod go.sum ./
RUN go mod tidy
COPY . ./ 
RUN go build -o /app/server cmd/server/main.go

# セキュリティ的にはshellが含まれていないgcr.io/distroless/static-debian12などが良いが、デバッグの容易性のためにalpineを選択した。
# セキュリティ要件によっては上記のイメージを選択してもいいだろう。
FROM alpine
RUN apk --no-cache add tzdata
COPY --from=builder /app/server /server
USER 1001
CMD ["/server"]
