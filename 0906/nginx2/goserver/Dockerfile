FROM golang:alpine

RUN apk add git

WORKDIR /go/src/server
RUN go get -u github.com/golang/dep/...
COPY src/ .
RUN dep ensure
RUN go install

CMD /go/bin/server
