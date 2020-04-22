FROM golang:1.14

ENV GO111MODULE=on

# Install development dependancies
RUN go get -tags nomysql github.com/steinbacher/goose/cmd/goose

RUN mkdir -p /go/src/github.com/deanobarnett/mood-tracker
WORKDIR /go/src/github.com/deanobarnett/mood-tracker
COPY . .

RUN go mod verify
RUN GOFLAGS=-mod=vendor go install -race ./cmd/...

CMD ["/go/bin/mood-tracker"]
