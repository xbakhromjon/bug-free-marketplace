#!/bin/bash
CURRENT_DIR=$(pwd)

source "$CURRENT_DIR"/scripts/parse_yaml.sh
# shellcheck disable=SC2046
eval $(parse_yaml genprotos.yaml "genprotos_")

rm -rf ./genproto/*

for pkg in ${genprotos_proto_list[*]}; do
  for module in $(find $CURRENT_DIR/iman_protos/$pkg -type d); do
      protoc -I=${module} -I $CURRENT_DIR/iman_protos/ \
             --gofast_out=plugins=grpc:$CURRENT_DIR/ \
              $module/*.proto;
  done;
done;
