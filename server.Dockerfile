# workspace (GOPATH) configured at /go
FROM golang:1.20.7 as builder


RUN mkdir app
WORKDIR app

COPY  go.mod .
COPY  go.sum .
RUN go mod download -x

COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make server-build && \
    mv ./bin/server /


FROM alpine
COPY --from=builder server .
COPY .env .env
COPY static/quotes.txt static/quotes.txt

ENTRYPOINT ["./server"]
