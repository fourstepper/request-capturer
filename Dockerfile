# Start from the official Golang base image
FROM golang:1.22-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
# Uncomment the lines below if you are using go.mod and go.sum files for dependencies
COPY go.mod ./

# Download any dependencies
# Uncomment the line below to enable modules download if you are using go.mod and go.sum
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY main.go .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
