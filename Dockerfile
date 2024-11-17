FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./main ./main.go

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8000
ENTRYPOINT ["./main"]