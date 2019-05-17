#!/bin/bash

if ! [ -x "$(command -v go)" ]; then
  echo 'Error: go is not installed. You can download it from: https://golang.org/dl/.' >&2
  exit 1
fi

present=$(go env GOPATH)/bin/present
if ! [ -x "$present" ]; then
  echo 'Error: go presenter is not installed. Install with: go install golang.org/x/tools/cmd/present.' >&2
  exit 1
fi

$present -orighost localhost
