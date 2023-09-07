
FROM golang:1.19.2-alpine

WORKDIR /app

COPY . /app

<<<<<<< HEAD
=======
# Fixed typo in the command. Changed "go build -o app ." to "go build -o main ."
>>>>>>> 0b9e1dfca7478780492ca6af66f157ed6fe805f2
RUN go build -o main .

EXPOSE 8080

# Added a comment explaining the purpose of the exposed port
# Changed the CMD to use the correct binary name "main" instead of "app/main"
CMD ["/app/main"]
