FROM golang:1.24.0 AS builder

WORKDIR /app

COPY . .

RUN go build -o main

FROM golang:1.24.0-slim

WORKDIR /app

COPY --from=builder /app/main .

CMD [ "./main" ]
