FROM golang:1.19 AS builder

ARG VERSION=dev

WORKDIR /go/src/app
COPY . .
RUN go build -o main -ldflags=-X=min.version=${VERSION} main.go

FROM debian:buster-slim
COPY --from=builder /go/src/app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
CMD [ "main" ]