#!/bin/bash

set -e

version=v$(cat VERSION)

builds=(
	"dummyfossilizer github.com/stratumn/go/cmd/dummyfossilizer"
	"dummystore github.com/stratumn/go/cmd/dummystore"
	"filestore github.com/stratumn/go/cmd/filestore"
)

targets=(
#	"darwin 386"
	"darwin amd64"
#	"darwin arm"
#	"darwin arm64"
#	"dragonfly amd64"
#	"freebsd 386"
#	"freebsd amd64"
#	"freebsd arm"
#	"linux 386"
	"linux amd64"
#	"linux arm"
#	"linux arm64"
#	"linux ppc64"
#	"linux ppc64le"
#	"linux mips64"
#	"linux mips64le"
#	"netbsd 386"
#	"netbsd amd64"
#	"netbsd arm"
#	"openbsd 386"
#	"openbsd amd64"
#	"openbsd arm"
#	"plan9 386"
#	"plan9 amd64"
#	"solaris amd64"
#	"windows 386"
	"windows amd64"
)

for (( i = 0; i < ${#targets[@]}; i++ )); do
	target=(${targets[i]})
	os=${target[0]}
	arch=${target[1]}

	for (( j = 0; j < ${#builds[@]}; j++ )); do
		build=(${builds[j]})
		name=${build[0]}
		package=${build[1]}

		dir=dist/${os}-${arch}/${name}

		if [[ $os == "windows" ]]; then
			name=${name}.exe
		fi

		mkdir -p $dir

		out=${dir}/${name}

		echo Building $out
		GOOS=$os GOARCH=$arch go build -o "$(pwd)/${out}" -ldflags "-X main.version=${version}" "$package"
	done
done
