# Stage 1: Modules caching
FROM golang:1.24.4-alpine AS modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Stage 2: Builder
FROM golang:1.24.4-alpine AS builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/CarRent ./cmd/main.go

# Stage 3: Final
FROM alpine:latest
COPY --from=builder /bin/CarRent /app/CarRent
COPY --from=builder /app/configs /app/configs
WORKDIR /app
EXPOSE ${HTTP_PORT}
CMD ["/app/CarRent"]