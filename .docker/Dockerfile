FROM golang:1.22.0-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/contained

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main /

EXPOSE 8080

CMD ["/main"]
