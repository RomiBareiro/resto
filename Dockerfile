FROM golang:1.22 AS builder

RUN apt-get update && apt-get install -y \
    gcc \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /resto_go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/template/csv_info.csv /app/template/csv_info.csv

COPY --from=builder /resto_go /app/resto_go

EXPOSE 8080

ENTRYPOINT ["/app/resto_go"]
