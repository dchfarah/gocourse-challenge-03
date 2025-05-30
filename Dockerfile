FROM golang:1.24.3-alpine3.21 AS builder
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o ordersystem ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

FROM scratch
COPY --from=builder /app/ordersystem /app/cmd/ordersystem/.env .
CMD ["./ordersystem"]
