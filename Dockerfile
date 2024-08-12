FROM golang:1.22.6-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /app/server cmd/server/main.go

# セキュリティ的にはshellが含まれていないgcr.io/distroless/static-debian12などが良いが、デバッグの容易性のためにalpineを選択した。
# セキュリティ要件によっては上記のイメージを選択してもいいだろう。
FROM alpine
RUN apk --no-cache add tzdata
COPY --from=builder /app/server /server
USER 1001
CMD ["/server"]
