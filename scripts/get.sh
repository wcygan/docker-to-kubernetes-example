#!/bin/bash

if [ $# -eq 1 ]
then
  go run client/main.go get $1
elif [ $# -eq 0 ]
then
  go run client/main.go get foo
else
  echo "Invalid number of arguments. Expected 0 or 1, got $#"
fi