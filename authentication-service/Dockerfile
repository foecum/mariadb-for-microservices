FROM golang:1.11.5-alpine as builder

ADD . /go/src/mariadb.com/authentication-service/
WORKDIR /go/src/mariadb.com/authentication-service
RUN go build .

FROM alpine:3.8
COPY --from=builder /go/src/mariadb.com/authentication-service/authentication-service .
ENTRYPOINT ["./authentication-service"]