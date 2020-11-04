FROM golang:1.14-alpine AS builder

WORKDIR /builder
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/blog/

FROM alpine:3.12.1

WORKDIR /app
RUN apk add --no-cache \
    ca-certificates \
    tzdata && \
    cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && \
    echo "Asia/Seoul" > /etc/timezone

COPY --from=builder /builder/server /app/program/server
COPY --from=builder /builder/scripts/entrypoint.sh /app/program/entrypoint.sh

RUN chmod +x /app/program/server
RUN chmod +x /app/program/entrypoint.sh

ENTRYPOINT [ "/app/program/entrypoint.sh" ]


