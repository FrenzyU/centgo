# Use the official Go image from the Docker Hub
FROM golang:1.17-alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Create a directory in the container to hold the source code
WORKDIR /app

# Copy the source code into the container
COPY . .

# Fetch the dependencies using go get
RUN go get -d -v

# Build the application
RUN go build -o /main .

# Expose port 8080 to the outside
EXPOSE 8080

# Command to run when starting the container
CMD ["/main"]
