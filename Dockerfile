FROM golang:1.19

# Unsetting GOPATH as it is unnecessary when working with modules
WORKDIR /go/src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod .
RUN go mod verify

COPY . .
RUN go build -v -o ./gt-branch-remover ./...

ENTRYPOINT [ "gt-branch-remover" ]