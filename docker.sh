#!/bin/bash
set -e

version=$(cat VERSION)
cd dist/linux-amd64

for cmd in ./*; do
	[ -d "${cmd}" ] || continue
	cd $cmd
	cmd=$(basename $cmd)
	tee Dockerfile > /dev/null <<EOF
FROM alpine:3.3

RUN mkdir -p /opt/stratumn/bin
ADD $cmd /opt/stratumn/bin/
ADD LICENSE /opt/stratumn/
WORKDIR /opt/stratumn/bin/

CMD ["./${cmd}"]
EOF
	docker build -t stratumn/${cmd}:${version} .
	cd ..
done

cd ..
