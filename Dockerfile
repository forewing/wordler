FROM golang:1.17-alpine AS builder

WORKDIR /build
COPY . /build/
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath -o wordler-web ./cmd/wordler-web

FROM alpine:3
RUN apk add --no-cache \
    dumb-init

WORKDIR /app
COPY --from=builder /build/wordler-web /app/

EXPOSE 8080
ENTRYPOINT [ "dumb-init", "--", "/app/wordler-web" ]
