FROM golang:1.19

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN ls
RUN pwd
RUN go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app" "tarakaprabhuchinta" "azure-infra-flux-apps"]