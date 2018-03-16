#!/bin/bash

echo "Running 'dep ensure'..."
dep ensure
echo "Running 'dep ensure'... DONE!"

for D in cmd/*; do
    if [ -f ${D}/main.go ]; then
        PACKAGE=$(echo ${D##*/})
        echo "Building $PACKAGE..."
        GOOS=linux go build -ldflags="-s -w" -o bin/${PACKAGE} ${D}/main.go
        echo "Building $PACKAGE... DONE!"
    fi
done
