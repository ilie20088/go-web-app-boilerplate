#!/bin/sh
files=$(find . -not -path './vendor*' -name '*.go')

# go fmt
unfmtd=$(echo $files | xargs gofmt -l)
if [[ ! -z $unfmtd ]]; then
    	echo "Some .go files aren't formatted:\n$unfmtd"
fi

# go vet
echo "Running go vet checks..."
echo $files | xargs go tool vet -all

# golint
unlinted=$(golint $(glide novendor))
if [[ ! -z $unlinted ]]; then
	echo "golint issues found:\n$unlinted"
fi

exit 0