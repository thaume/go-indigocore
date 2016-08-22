FROM alpine:latest

RUN mkdir -p /opt/stratumn/bin
ADD linux_amd64/{{CMD}} /opt/stratumn/bin/
WORKDIR /opt/stratumn/bin/

CMD ["./{{CMD}}"]
