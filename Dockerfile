# Builder Stage
FROM golang:1.17.2-alpine3.14 as builder

# Copy application to the container
COPY . /go/src/url-short

# Get dependencies
WORKDIR /go/src/url-short
RUN go get -d -v ./...

# Build the application
WORKDIR /go/src/url-short/bin
RUN go build -o ./url-short ../cmd/url-short/main.go

# Run Stage
FROM alpine:3.9

COPY --from=builder /go/src/url-short/bin /app/url-short

ENV URL_SHORT_PORT="8000"
ENV URL_SHORT_REDIS_PASSWORD="zxcv1234"

EXPOSE 8000

# Run application
WORKDIR /app/url-short
CMD ["./url-short"]
