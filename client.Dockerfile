# workspace (GOPATH) configured at /go
FROM golang:1.20.7 as builder


RUN mkdir app
WORKDIR app

COPY  go.mod  .
COPY  go.sum .
RUN go mod download -x

COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make client-build && \
    mv ./bin/client /

FROM alpine
COPY --from=builder client .
COPY .env .env

ENTRYPOINT ["./client"]
