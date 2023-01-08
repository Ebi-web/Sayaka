FROM golang:1.19

# Set the working directory to the project root
WORKDIR /app

# Copy the application source code
COPY ./app .

RUN go mod tidy

# Build the Go application
RUN go build -o main .

# Expose the application's port
EXPOSE 8080

# Set the entrypoint to the application's binary
ENTRYPOINT ["./main"]
