#!/bin/bash
set -e

cd dist

for target in ./*; do
	[ -d "${target}" ] || continue

	cd $target

	for cmd in ./*; do
		[ -d "${cmd}" ] || continue

		cp ../../LICENSE $cmd/
		cp ../../RELEASE_NOTES.md $cmd/
		cp ../../CHANGE_LOG.md $cmd/
		zip -r ${cmd}{.zip,}
	done

	cd ..
done

cd ..
