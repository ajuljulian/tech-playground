# Set the base image of our image
FROM golang:1.17

# Create a /app directory inside our container that will hold the
# application source files
RUN mkdir /app

# Copy everything in the root directory into /app
ADD . /app

# Specify that we want to execute any further commands inside
# our /app directory
WORKDIR /app

# Download all the dependencies of the main module into the module cache
# https://go.dev/ref/mod#go-mod-download
RUN go mod download

# Compile the code
RUN go build -o main .

# Start the executable
CMD ["/app/main"]
