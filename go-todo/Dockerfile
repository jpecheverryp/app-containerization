# Specifies a parent image
FROM golang:latest

# Creates an app directory to hold your app's source code
WORKDIR /app

# Copies everything from root directory into app directory
COPY . .

# Install Golang Dependencies
RUN go mod download

# Build app
RUN go build -o /godocker

# Expose port to outside the container
EXPOSE 8080

# Specify the executable file that runs when the container starts
CMD [ "/godocker" ]
