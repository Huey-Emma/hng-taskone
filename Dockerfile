FROM golang:1.19.2-alpine

WORKDIR /app

COPY . /app

RUN go build -o app.go .

EXPOSE 8080

CMD ["/app/main"]
