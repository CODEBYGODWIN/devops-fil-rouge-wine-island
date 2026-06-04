# syntax=docker/dockerfile:1

FROM golang:1.25.3-alpine3.22 AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /bin/api-go .

FROM alpine:3.22

RUN addgroup -S app && adduser -S -G app app

WORKDIR /app

ENV PORT=3000

COPY --from=builder /bin/api-go ./api-go

USER app

EXPOSE 3000

CMD ["./api-go"]
