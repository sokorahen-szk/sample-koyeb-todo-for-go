# See https://www.koyeb.com/docs/deploy/go
ARG WORKING_DIR=/app
ARG APP_NAME=main
ARG APP_PORT=8000

FROM golang:1.23-alpine AS builder
WORKDIR ${WORKING_DIR}
COPY . .
RUN go mod download
RUN go build -o ./${APP_NAME} main.go

FROM alpine:latest AS runner
WORKDIR ${WORKING_DIR}
COPY --from=builder ${WORKING_DIR}/${APP_NAME} .
EXPOSE ${APP_PORT}
ENTRYPOINT ["./${APP_NAME}"]