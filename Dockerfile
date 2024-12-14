FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o app ./cmd/app

RUN chmod +x app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

CMD [ "/app/app" ]

