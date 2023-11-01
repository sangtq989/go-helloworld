# Build stage
FROM golang:1.16 as builder
WORKDIR /app
COPY main.go go.* /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 8080

# Run the Go binary in the runtime stage
CMD ["/app/main"]
