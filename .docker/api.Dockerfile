FROM golang:1.19.1-buster

ENV ROOT=/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN go install github.com/cosmtrek/air@latest

RUN apt update && apt-get install git
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o ./main ./src/cmd/main.go

CMD ["air", "-c", ".air.toml"]
