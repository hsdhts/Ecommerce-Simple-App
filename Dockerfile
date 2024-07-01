# Use an official Golang runtime as a parent image
FROM golang:1.16.3-alpine3.13

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 4000 to the outside world
EXPOSE 4000

# Run the executable
CMD ["./main"]
