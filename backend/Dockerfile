FROM golang:1.22-alpine AS builder
RUN apk add --no-cache gcc musl-dev sqlite-dev
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o task-manager-backend main.go

FROM alpine:3.18
WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/task-manager-backend .

ENV PORT=8080
EXPOSE 8080

CMD ["./task-manager-backend"]