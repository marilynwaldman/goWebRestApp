FROM golang:latest


RUN go get github.com/julienschmidt/httprouter
RUN go get 	github.com/rs/cors

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/user/goWebRestApp/api

# Install our dependencies


# Install api binary globally within container

RUN go install github.com/user/goWebRestApp/api/src


# Set binary as entrypoint
ENTRYPOINT /go/bin/src

# Expose default port (3000)
EXPOSE 3000