FROM golang:1.19

COPY . .
RUN go build -v -o gt-branch-remover ./...
ENTRYPOINT [ "./gt-branch-remover" ]
