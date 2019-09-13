FROM golang:1.12-alpine
MAINTAINER PeterSamokhin https://github.com/petersamokhin

RUN apk add --no-cache git wget
RUN go version
RUN go get golang.org/x/tools/cmd/godoc
RUN which godoc && which wget

RUN mkdir -p /app/docs
RUN mkdir -p /app/gomodgodoc
WORKDIR /app

COPY go.mod go.sum ./
COPY gomodgodoc/go.mod ./gomodgodoc
RUN go mod download

COPY main.go replacer.json .env ./
COPY gomodgodoc/replacer.go gomodgodoc/utils.go gomodgodoc/gomodgodoc.go ./gomodgodoc/

RUN go build -o main .

CMD ["./main"]
