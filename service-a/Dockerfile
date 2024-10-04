# Build stage
FROM golang:1.16-alpine AS build

# Install protobuf compiler
RUN apk add --no-cache protobuf protobuf-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Generate protobuf files
RUN protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative proto/service.proto

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Run stage
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]