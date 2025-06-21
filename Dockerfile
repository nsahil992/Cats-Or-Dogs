# ----- BUILD STAGE -----
FROM golang:1.22 AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/ ./

RUN GOOS=linux GOARCH=amd64 go build -o main

# ----- RUN STAGE -----
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .env
COPY frontend ./frontend

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]
