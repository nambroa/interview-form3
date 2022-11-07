FROM golang:1.19 as development

# Create dir for app
WORKDIR /form3-interview-app

# Copy app Dependencies (validator, testify, etc..)
COPY go.mod go.sum ./
# Run command to actually download the dependencies
RUN go mod download

# Copy interview app to docker container.
COPY . .


# TODO: REMOVE
# Build the go app and create the binary
# TODO: Maybe here a command to run the tests????
#RUN go test ./...

# Expose the port to talk to the app and run the app
#EXPOSE 4000
#CMD [ "/form3-interview-app" ]
