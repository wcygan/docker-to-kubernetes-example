#!/bin/bash

if [ $# -eq 2 ]
then
  grpcurl -plaintext -proto proto/kv/v1/kv.proto -d '{ "key": "'"$1"'", "value": "'"$2"'" }' localhost:8080 kv.v1.KeyValueService/Put
elif [ $# -eq 0 ]
then
  grpcurl -plaintext -proto proto/kv/v1/kv.proto -d '{ "key": "foo", "value": "bar" }' localhost:8080 kv.v1.KeyValueService/Put
else
  echo "Invalid number of arguments. Expected 0 or 2, got $#"
fi