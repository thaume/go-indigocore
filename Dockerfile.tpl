FROM alpine:latest

RUN apk update && apk upgrade
RUN apk add --no-cache ca-certificates
RUN mkdir -p /usr/local/stratumn/bin
COPY LICENSE /usr/local/stratumn/
COPY dist/linux-amd64/{{CMD}} /usr/local/stratumn/bin/
RUN mkdir -p /usr/local/bin
RUN ln -s /usr/local/stratumn/bin/* /usr/local/bin

CMD ["{{CMD}}"]
