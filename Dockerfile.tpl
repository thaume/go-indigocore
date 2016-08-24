FROM alpine:latest

RUN apk add --no-cache ca-certificates
RUN mkdir -p /opt/stratumn/bin
ADD LICENSE /opt/stratumn/
ADD dist/linux-amd64/{{CMD}} /opt/stratumn/bin/
WORKDIR /opt/stratumn/bin/

CMD ["./{{CMD}}"]
