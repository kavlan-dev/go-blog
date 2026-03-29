FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

FROM alpine
WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]
