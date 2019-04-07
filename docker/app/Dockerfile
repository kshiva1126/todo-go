FROM golang:1.12

WORKDIR /go/src/todo-go
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]