#!/bin/bash
set -e


echo "==> Running tests"
./coverage.sh

echo "==> Running linter"
golint -set_exit_status ./...
