FROM golang:1.24.0-alpine3.21

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main ./cmd

EXPOSE 7000

CMD ["./main"]
