FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o observ .

FROM scratch

COPY --from=builder /app/observ /observ

USER 1000:1000

EXPOSE 8080

ENTRYPOINT ["/observ"]