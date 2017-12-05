FROM golang:latest


RUN go get github.com/julienschmidt/httprouter
RUN go get 	github.com/rs/cors

run echo
# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/user/goWebRestApp/api

# Install ourcd  dependencies

# Install api binary globally within container
RUN go install github.com/user/goWebRestApp/api/models
RUN go install github.com/user/goWebRestApp/api/controllers
RUN go install github.com/user/goWebRestApp/api/src
WORKDIR /go/src/github.com/user/goWebRestApp/api


# Set binary as entrypoint
ENTRYPOINT /go/bin/src

# Expose default port (3000)
EXPOSE 3000