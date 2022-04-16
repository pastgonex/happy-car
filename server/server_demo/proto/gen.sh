#! /bin/zsh
protoc -I. --go_out=paths=source_relative:gen/go --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:gen/go trip.proto
protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=trip.yaml:gen/go trip.proto

PBTS_BIN=../../wx/miniprogram/node_modules/.bin
PBTS_OUT_DIR=../../wx/miniprogram/service/proto_gen

${PBTS_BIN}/pbjs -t static -w es6 trip.proto --no-create --no-encode --no-decode --no-verify --no-delimited > ${PBTS_OUT_DIR}/trip_pb_temp.js
echo 'import * as $protobuf from "protobufjs";\n' > ${PBTS_OUT_DIR}/trip_pb.js
cat ${PBTS_OUT_DIR}/trip_pb_temp.js >> ${PBTS_OUT_DIR}/trip_pb.js
rm ${PBTS_OUT_DIR}/trip_pb_temp.js

${PBTS_BIN}/pbts -o ${PBTS_OUT_DIR}/trip_pb.d.ts ${PBTS_OUT_DIR}/trip_pb.js 

