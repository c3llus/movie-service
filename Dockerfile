# Use the official Golang image as a parent image
FROM golang:1.18

# Set the working directory to /app
WORKDIR /movie-service

# Copy the current directory contents into the container at /app
COPY . .

# Run `go mod download` to download and cache Go modules
RUN go mod download

# Build the Go program
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Define the command to run the executable
CMD ["./main"] 