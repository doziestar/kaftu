FROM golang:1.19.0-alpine3.13 AS builder

WORKDIR /app/

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server

FROM alpine:3.13 AS runner

WORKDIR /app/

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]