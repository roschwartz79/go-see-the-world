# Stage 1: Build the Go application
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /go-see-the-world

# Stage 2: Create the final, minimal image
FROM alpine:latest

WORKDIR /

COPY --from=builder /go-see-the-world .

EXPOSE 8080

CMD ["./go-see-the-world"]