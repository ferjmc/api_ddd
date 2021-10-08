# Build stage
FROM golang:1.16-alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY go.mod . 
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o session cmd/main.go

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/session .

EXPOSE 5000
CMD ["./session"]