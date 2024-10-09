FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main ./cmd/api/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/main ./main
CMD ["/app/main"]