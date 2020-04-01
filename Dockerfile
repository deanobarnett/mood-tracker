FROM golang:1.14

# Install tools to make the container useful
RUN apt-get update
RUN apt-get install --yes vim

# Install some development dependancies
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.24.0
RUN go get -tags nomysql github.com/steinbacher/goose/cmd/goose

RUN mkdir -p /go/src/github.com/deanobarnett/mood-tracker
WORKDIR /go/src/github.com/deanobarnett/mood-tracker

ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go install -v ./cmd/...

CMD ["/go/bin/mood-tracker"]
