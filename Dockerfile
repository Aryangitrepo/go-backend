FROM golang:1.22.2

WORKDIR /usr/src/app

# Copy the .env file for environment variables
COPY .env ./

# Copy your compiled Go binary (named 'main') into the container
COPY app ./

# Run the Go binary
CMD ["./main"]
