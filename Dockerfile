# Use the official Go image as the base
FROM golang:latest

# hot reload
RUN go install github.com/cosmtrek/air@latest

RUN mkdir /app

# Copy the rest of the application code
ADD . /app

# Set the working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 9999

# Command to run the application
# CMD ["go", "run", "main.go"]
ENTRYPOINT [ "air" ]
