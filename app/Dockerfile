FROM golang:1.19

# Set the working directory to the project root
WORKDIR /app

# Copy the application source code
COPY . .

RUN go mod tidy

# Expose the application's port
EXPOSE 8080

RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
