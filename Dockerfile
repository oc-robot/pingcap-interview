FROM golang:1.15

WORKDIR /root

COPY . /root

RUN go build -o /bin/demo cmd/main.go

EXPOSE 2332

ENTRYPOINT ["/bin/demo"]
