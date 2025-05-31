# Build stage
FROM golang:1.24.3 AS builder

WORKDIR /app
COPY . .

# Build statically-linked binary
ENV CGO_ENABLED=0
RUN go build -o server ./cmd/server

# Run stage using scratch (or use alpine for debugging)
FROM alpine:3.20

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/server /server

EXPOSE 9091
ENTRYPOINT ["/server"]
