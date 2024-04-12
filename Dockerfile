FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk --no-cache add bash git make gettext

COPY go.* ./
RUN go mod download

COPY TestTask-PGPro ./

RUN go build -o ./bin/launch-service cmd/main.go

FROM alpine AS runner

COPY --from=builder /app/bin/launch-service /
COPY --from=builder /app/migrations /migrations
COPY internal/config/dev_config.json config.json

ENTRYPOINT ["./launch-service", "-path=/config.json"]