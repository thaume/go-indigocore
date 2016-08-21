FROM golang:1.7-alpine

RUN apk update
RUN apk upgrade
RUN apk add --no-cache git

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
