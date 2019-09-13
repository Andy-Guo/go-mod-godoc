FROM golang:1.12-alpine
MAINTAINER PeterSamokhin https://github.com/petersamokhin

RUN apk add --no-cache git wget
RUN go version
RUN go get golang.org/x/tools/cmd/godoc
RUN which godoc && which wget

RUN mkdir -p /app/docs
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go replacer.go utils.go replacer.json .env ./

RUN go build -o main .

CMD ["./main"]
