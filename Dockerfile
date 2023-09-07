
FROM golang:1.19.2-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the application using the correct binary name: main
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 8080

# Set the command to run when the container starts
CMD ["/app/main"]
