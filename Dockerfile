FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/app

FROM alpine AS final
WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]
