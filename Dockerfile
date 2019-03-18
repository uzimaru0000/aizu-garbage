FROM golang:latest as builder
COPY ./ /go/src/github.com/uzimaru0000/aizu-garbage
WORKDIR /go/src/github.com/uzimaru0000/aizu-garbage
RUN go get -u github.com/golang/dep/cmd/dep && dep ensure && make

FROM alpine
ENV DOCKERIZE_VERSION v0.6.0
COPY --from=builder /go/src/github.com/uzimaru0000/aizu-garbage/api ./
COPY --from=builder /go/src/github.com/uzimaru0000/aizu-garbage/migrator ./
RUN apk add --no-cache openssl &&\
    wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz &&\
    tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz &&\
    rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
CMD [ "/migrator", "&&", "/api" ]
