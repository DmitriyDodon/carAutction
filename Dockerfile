# Use the official golang image as a base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /usr/local/go/src/app

RUN touch data.db

# Copy the local package files to the container's workspace
COPY . .

RUN wget https://github.com/swaggo/swag/releases/download/v1.16.3/swag_1.16.3_Linux_amd64.tar.gz

RUN tar -xvzf swag_1.16.3_Linux_amd64.tar.gz

RUN go mod vendor

RUN ./swag init -g ./cmd/main.go

# Download and install any required dependencies
RUN go mod download -x

# Build the Go app
RUN go build -o main ./cmd

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]