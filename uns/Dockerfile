# Use the official golang image as the base image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Download and install the required dependencies
RUN go get -d -v ./...

# Build the Go application
RUN go build -o main .

# Expose port 8050 to the outside world
EXPOSE 8050

# Run the binary
CMD ["./main"]
