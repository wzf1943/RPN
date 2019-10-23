FROM golang:1.10 AS builder

COPY . /go/src/RPN

WORKDIR /go/src/RPN

RUN go get github.com/gorilla/mux && go get github.com/golang-collections/collections/stack && go get github.com/jessevdk/go-flags

RUN go build -o rpn -v

FROM oraclelinux:7-slim

COPY --from=builder /go/src/RPN/rpn /bin

CMD ["rpn"]