# ---------- build stage ----------
FROM golang:1.26.1-alpine AS builder

WORKDIR /src

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/app .

# ---------- runtime stage ----------
FROM alpine:3.20

RUN apk add --no-cache ca-certificates && \
    adduser -D -u 10001 appuser

WORKDIR /app
COPY --from=builder /out/app /app/app

USER appuser
EXPOSE 8080

ENV GIN_MODE=release
ENV AUTHOR=""

ENTRYPOINT ["/app/app"]