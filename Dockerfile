
FROM golang:1.19.2-alpine

WORKDIR /app

COPY . /app

# Fixed typo in the command. Changed "go build -o app ." to "go build -o main ."
RUN go build -o main .

EXPOSE 8080

# Added a comment explaining the purpose of the exposed port
# Changed the CMD to use the correct binary name "main" instead of "app/main"
CMD ["/app/main"]
