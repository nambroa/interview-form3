FROM golang:1.19 as development

# Create dir for app
WORKDIR /form3-interview-app

# Copy app Dependencies (validator, testify, etc..)
COPY go.mod go.sum ./
# Run command to actually download the dependencies
RUN go mod download

# Copy interview app to docker container.
COPY . .
