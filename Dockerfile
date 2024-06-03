# Use an appropriate Golang base image
FROM golang:alpine AS build-env

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files for dependency management
COPY go.mod ./
COPY go.sum ./

# Fetch dependencies
RUN go mod download

# Copy the project source code
COPY . ./

# Build the application within the container 
RUN go build -o main ./cmd/main

# Define a more minimal runtime image
FROM alpine

# Set the workdir for the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=build-env /app/main /app/main

# Configure the default command to run when the container starts
CMD [ "/app/main" ]
