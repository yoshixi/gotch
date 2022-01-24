FROM golang:1.17-alpine
RUN apk update
RUN mkdir -p /go/src/github.com/yoshixj/gotch
WORKDIR /go/src/github.com/yoshixj/gotch
ADD . /go/src/github.com/yoshixj/gotch

RUN go mod download
RUN go build -o gotch

