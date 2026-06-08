FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o hello-api .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /build/hello-api .

EXPOSE 9090

CMD ["./hello-api"]
