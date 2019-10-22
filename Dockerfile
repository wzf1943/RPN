FROM golang:1.10 AS builder

COPY . /go/src/RPN

WORKDIR /go/src/RPN

RUN go get github.com/gorilla/mux && go get github.com/golang-collections/collections/stack

RUN go build

FROM oraclelinux:7-slim

COPY --from=builder /go/src/RPN/RPN /bin

CMD ["RPN"]