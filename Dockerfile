FROM golang:1.19

COPY main.go ./main.go
COPY json-decode.go ./json-decode.go
RUN go build -v -o gt-branch-remover ./...
ENTRYPOINT [ "./gt-branch-remover" ]
