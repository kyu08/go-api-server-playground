FROM golang:1.22.6-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /app/server cmd/server/main.go

FROM alpine
RUN apk --no-cache add tzdata
COPY --from=builder /app/server /server
USER 1001
CMD ["/server"]
