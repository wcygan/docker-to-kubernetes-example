#!/bin/bash

if [ $# -eq 2 ]
then
  go run client/main.go put $1 $2
elif [ $# -eq 0 ]
then
  go run client/main.go put foo bar
else
  echo "Invalid number of arguments. Expected 0 or 2, got $#"
fi