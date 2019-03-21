FROM golang:1.12

WORKDIR /go/src/hot_reload_docker
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]