#!/bin/bash

#!/bin/bash

if [ $# -eq 1 ]
then
  grpcurl -plaintext -proto proto/kv/v1/kv.proto -d '{ "key": "'"$1"'" }' localhost:8080 kv.v1.KeyValueService/Get
elif [ $# -eq 0 ]
then
  grpcurl -plaintext -proto proto/kv/v1/kv.proto -d '{ "key": "foo" }' localhost:8080 kv.v1.KeyValueService/Get
else
  echo "Invalid number of arguments. Expected 0 or 1, got $#"
fi