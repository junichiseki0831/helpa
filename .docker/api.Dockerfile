FROM golang:1.19.1-buster

ENV CGO_ENABLED 0
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

RUN apt update && apt install git
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./

CMD ["air", "-c", ".air.toml"]
