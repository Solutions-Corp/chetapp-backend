FROM golang:1.24-alpine@sha256:ef18ee7117463ac1055f5a370ed18b8750f01589f13ea0b48642f5792b234044 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o routes-service ./cmd/main.go

FROM alpine:latest@sha256:a8560b36e8b8210634f77d9f7f9efd7ffa463e380b75e2e74aff4511df3ef88c

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/routes-service /routes-service

EXPOSE 8082

CMD ["/routes-service"]