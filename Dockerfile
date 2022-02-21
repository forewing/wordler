FROM golang:1.17-alpine AS builder
RUN apk add --no-cache \
    git

WORKDIR /build
COPY . /build/
RUN go run ./cmd/build -web

FROM alpine:3
RUN apk add --no-cache \
    dumb-init

WORKDIR /app
COPY --from=builder /build/output /app/

EXPOSE 8080
ENTRYPOINT [ "dumb-init", "--", "/app/wordler-web" ]
