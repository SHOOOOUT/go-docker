FROM golang:1.15.2-alpine3.12

WORKDIR /Users/shuto/code/golang/src/github.com/shuto/go-docker
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]