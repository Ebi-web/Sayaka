FROM golang:1.19

# Set the working directory to the project root
WORKDIR /app

# Copy the application source code
COPY . .

# Expose the application's port
EXPOSE 8080

RUN go install github.com/cosmtrek/air@latest && \
    go mod tidy
CMD ["air"]
