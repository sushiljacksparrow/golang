FROM golang:1.13-alpine
EXPOSE 8080

ENV GOPATH=/go
RUN mkdir -p $GOPATH/src/github.com/rest-go
COPY . $GOPATH/src/github.com/rest-go

WORKDIR $GOPATH/src/github.com/rest-go
RUN go build -o rest-go .

CMD ["/go/src/github.com/rest-go/rest-go"]
