FROM golang:1.11.5-alpine as builder

ADD . /go/src/mariadb.com/comment-service/
WORKDIR /go/src/mariadb.com/comment-service
RUN go build .

FROM alpine:3.8
COPY --from=builder /go/src/mariadb.com/comment-service/comment-service /
CMD ["/comment-service"]