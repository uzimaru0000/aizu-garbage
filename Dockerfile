FROM golang:latest as builder
COPY ./ /go/src/github.com/uzimaru0000/aizu-garbage
WORKDIR /go/src/github.com/uzimaru0000/aizu-garbage
RUN make

FROM alpine
ENV DOCKERIZE_VERSION v0.6.0
COPY --from=builder /go/src/github.com/uzimaru0000/aizu-garbage/api ./
CMD [ "/api" ]
