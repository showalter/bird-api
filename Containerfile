# Create an environment for building the binary
FROM golang:1.15 AS build

WORKDIR /go/src/bird-api

# Copy all code for building
COPY . .

# Download dependencies
RUN go mod download

# Build staticly linked binary for alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


# Move binary to final image and run it
FROM alpine:latest

WORKDIR /app

# Copy built binary from build image
COPY --from=build go/src/bird-api/app .
COPY ./birds.json .

# Expose port 3000
EXPOSE 3000

CMD ["./app"]

