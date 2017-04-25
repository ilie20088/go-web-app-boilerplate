#!/bin/sh
# This script should be run before every commit to identify code formatting and style issues
files=$(find . -not -path './vendor*' -name '*.go')

# go fmt
unfmtd=$(echo $files | xargs gofmt -l)
if [[ ! -z $unfmtd ]]; then
    	echo "Some .go files aren't formatted:\n$unfmtd"
fi

# go vet
echo $files | xargs go tool vet -all

# golint
unlinted=$(golint $(glide novendor))
if [[ ! -z $unlinted ]]; then
	echo "golint issues found:\n$unlinted"
fi

exit 0