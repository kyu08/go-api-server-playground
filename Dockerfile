FROM golang:1.22.6-bookworm as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/server /app/cmd/server/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/server /server
USER 1001
CMD ["/server"]
