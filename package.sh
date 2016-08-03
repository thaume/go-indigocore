#!/bin/bash
set -e

cd dist

for target in ./*; do
	[ -d "${target}" ] || continue

	cd $target

	for package in ./*; do
		[ -d "${package}" ] || continue

		cp ../../LICENSE $package/
		cp ../../RELEASE_NOTES.md $package/
		cp ../../CHANGE_LOG.md $package/
		zip -r ${package}{.zip,}
	done

	cd ..
done

cd ..
