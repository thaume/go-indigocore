#!/bin/bash
set -e

commit=$(git rev-parse HEAD)
version=$(cat VERSION)
tag=v${version}
prerelease=$(cat PRERELEASE)

description=$(cat <<EOF
$(cat RELEASE_NOTES.md)

[![Build Status](https://travis-ci.org/stratumn/go.svg?branch=${tag})](https://travis-ci.org/stratumn/go)

[Change log](https://github.com/stratumn/go/blob/$commit/CHANGE_LOG.md)
EOF)

flags="--user stratumn --repo go --tag '$tag'"
release_flags="$flags --name 'Stratumn Go packages $tag' --description '$description'"

#./check.sh
#
#echo "==> Cleaning up"
#rm -rf dist
#
#echo "==> Building"
#./build.sh
#
#echo "==> Packaging"
#./package.sh
#
#echo "==> Creating docker images"
#./docker.sh
#
#echo "==> Creating tag $tag"
#git tag $tag 2>/dev/null || echo Tag $tag already exists
#
#echo "==> Pushing tags"
#git push origin --tags
#
#echo "==> Creating release"
#eval github-release release "$release_flags" --target "$commit" --draft
#
#echo "==> Uploading binaries"
#for target in ./dist/*; do
#	[ -d "${target}" ] || continue
#	for file in ${target}/*.zip; do
#		echo Uploading $file
#		eval github-release upload "$flags" --file "$file" --name "$(basename ${file/.zip/})-$(basename $target).zip"
#	done
#done

echo "==> Uploading docker images"
cd dist/linux-amd64
for cmd in ./*; do
	[ -d "${cmd}" ] || continue
	cd $cmd
	cmd=$(basename $cmd)
	docker push stratumn/${cmd}:${version}
	docker push stratumn/${cmd}:latest
	cd ..
done
cd ..
cd ..

echo "==> Publishing"
if [[ $prerelease != "false" ]]; then
	eval github-release edit "$release_flags" --pre-release
else
	eval github-release edit "$release_flags"
fi

echo "==> Done"
