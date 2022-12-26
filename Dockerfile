FROM golang:latest

# Add the source code
ADD . /app

# Set the working directory
WORKDIR /app

# Build the Go web server
RUN go build -o main .

# Expose the port for the Go web server
EXPOSE 8080

# Run the Go web server
CMD ["./main"]
