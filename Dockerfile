FROM golang:1.21-alpine as builder

LABEL stage=gobuilder

ENV CGO_ENABLED = 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .

RUN go mod download && go mod verify

RUN go build -ldflags="-s -w" -o /app/person cmd/main.go

FROM alpine

RUN apk update && apk upgrade

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

RUN adduser -D appuser
USER appuser

WORKDIR /app

COPY --from=builder /app/person /app/person
COPY --from=builder /build/.env /app

CMD ["./person"]

