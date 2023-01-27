FROM golang:1.19.1-alpine

ENV ROOT=/go/src
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./src ./
RUN go build src/core/cmd/main.go

CMD ["./main"]
