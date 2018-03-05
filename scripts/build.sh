#!/bin/bash

FILE_DIR=$(dirname "$0")
if [ $FILE_DIR = '.' ];
then
    PROJECT_DIR=$(dirname "$PWD")
else
    PROJECT_DIR=$(dirname "$PWD/$FILE_DIR")
fi
export GOPATH=$PROJECT_DIR/src && export GOBIN=$PROJECT_DIR/bin
echo "GOPATH: $GOPATH"
echo "GOBIN:  $GOBIN"
BINARY_FILE="permissionguard"
mkdir -p $PROJECT_DIR/bin
go build -o $BINARY_FILE $PROJECT_DIR/src/github.com/thyrlian/PermissionMatters/cmd/main.go
mv $BINARY_FILE $PROJECT_DIR/bin/$BINARY_FILE
echo -e "\nBinary file is generated to: $GOBIN/$BINARY_FILE"
